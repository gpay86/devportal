package com.gpay.gpaysample.GpaySampleApplication.clients;

import com.fasterxml.jackson.databind.JsonNode;
import com.google.gson.Gson;
import com.gpay.gpaysample.GpaySampleApplication.configs.Config;
import com.gpay.gpaysample.GpaySampleApplication.controller.response.GeneralResponse;
import com.gpay.gpaysample.GpaySampleApplication.utils.CryptoUtils;
import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;
import org.apache.commons.lang3.StringUtils;
import org.springframework.core.io.ClassPathResource;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;
import org.springframework.web.client.HttpClientErrorException;
import org.springframework.web.client.HttpServerErrorException;
import org.springframework.web.client.HttpStatusCodeException;
import org.springframework.web.client.RestClientResponseException;
import org.springframework.web.client.RestTemplate;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.security.GeneralSecurityException;
import java.security.KeyFactory;
import java.security.PrivateKey;
import java.security.Signature;
import java.security.spec.PKCS8EncodedKeySpec;
import java.time.Duration;
import java.time.Instant;
import java.util.Base64;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

@Component
@RequiredArgsConstructor
@Log4j2
public class BaseRestTemplate {
    private final RestTemplate restTemplate;
    private final Config config;
    private final Gson gsonSnackKey;
    private final Gson gson;
    private String token;

    public GeneralResponse request(String path, Object sampleRequest) {
        String baseRequestStr = gsonSnackKey.toJson(sampleRequest);
        HttpHeaders headers = buildHeader(baseRequestStr);
        HttpEntity<String> request = new HttpEntity<>(baseRequestStr, headers);
        GeneralResponse response = null;
        boolean isErorr = false;
        try {
            response = restTemplate.postForObject(config.getBaseUrl() + path, request, GeneralResponse.class);

        } catch (Exception e) {
            e.printStackTrace();
            response = new GeneralResponse();
            isErorr = true;
            if (e instanceof RestClientResponseException) {
                HttpStatusCodeException exT = (HttpStatusCodeException) e;
                if (HttpStatus.UNAUTHORIZED.equals(exT.getStatusCode())) {
                    clearToken();
                    // recall
                    headers = buildHeader(baseRequestStr);
                    request = new HttpEntity<>(baseRequestStr, headers);
                    response = restTemplate.postForObject(config.getBaseUrl() + path, request, GeneralResponse.class);
                } else {
                    String body = ((HttpServerErrorException) e).getResponseBodyAsString();
                    response = (gson.fromJson(body, GeneralResponse.class));
                }
            }
            log.error("api [{}] Request [{}] error [{}] ", path, baseRequestStr, e.getMessage());
        }
        if (!isErorr) {
            log.info("api [{}] Request [{}] response [{}] ", path, baseRequestStr, gson.toJson(response));
        }

        return response;
    }

    //3 minutes to reload cache
    @Scheduled(fixedRate = 3 * 60 * 1000)
    public void clearToken() {
        token = null;
    }

    private HttpHeaders buildHeader(Object request) {
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);
        String xTimestamp = System.currentTimeMillis() + "";
        String xRequestId = UUID.randomUUID().toString();
        headers.add("x-timestamp", xTimestamp);
        headers.add("x-requests-id", xRequestId);
        headers.add("x-certificate", rawPublicKey());
        headers.setBearerAuth(getToken());

        String rawData = xTimestamp + xRequestId + request;
        headers.add("signature", createSignature(rawData));
        return headers;
    }

    public String getToken() {
        if (StringUtils.isNotEmpty(token)) {
            return token;
        }
        HttpHeaders httpHeaders = new HttpHeaders();
        httpHeaders.setContentType(MediaType.APPLICATION_JSON);
        Map<String, String> param = new HashMap<>();
        param.put("client_id", config.getClientCode());
        param.put("client_secret", config.getClientSecret());
        HttpEntity request = new HttpEntity<>(param, httpHeaders);
        try {
            JsonNode response = restTemplate.postForObject(config.getBaseUrl() + "/auth/token", request, JsonNode.class);
            token = response.get("data").get("access_token").textValue();
            return token;
        } catch (Exception e) {
            e.printStackTrace();
            return "";
        }
    }

    private String createSignature(String data) {
        try {
            Signature privateSignature = Signature.getInstance("SHA256withRSA");
            privateSignature.initSign(privateKey());
            privateSignature.update(data.getBytes());
            byte[] signature = privateSignature.sign();

            return Base64.getEncoder().encodeToString(signature);
        } catch (Exception e) {
            e.printStackTrace();
            return null;
        }

    }

    public PrivateKey privateKey() throws IOException, GeneralSecurityException {
        String pathFile = new ClassPathResource(config.getPrivateKey()).getFile().getPath();
        byte[] keyBytes = Files.readAllBytes(Paths.get(pathFile));

        String key = new String(keyBytes, StandardCharsets.UTF_8);
        String privateKeyPem = key
                .replace("-----BEGIN PRIVATE KEY-----", "")
                .replace("-----END PRIVATE KEY-----", "")
                .replaceAll("\\s+", "");

        byte[] decoded = Base64.getDecoder().decode(privateKeyPem);
        PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(decoded);
        KeyFactory keyFactory = KeyFactory.getInstance("RSA");

        return keyFactory.generatePrivate(keySpec);
    }

    public String rawPublicKey() {
        try {
            String filePath = new ClassPathResource(config.getPublicKey()).getFile().getAbsolutePath();
            Path pathPem = Paths.get(filePath);
            byte[] keyBytes = Files.readAllBytes(pathPem);
            String temp = new String(keyBytes);
            String keyRaw = temp.replace("-----BEGIN CERTIFICATE-----", "")
                    .replace("-----END CERTIFICATE-----", "")
                    .replace("\n", "").replace("\r", "");
            return keyRaw;
        } catch (Exception e) {
            e.printStackTrace();
        }
        return null;
    }
}
