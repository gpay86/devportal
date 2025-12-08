package com.gpay.gpaysample.GpaySampleApplication.configs;

import com.google.gson.FieldNamingPolicy;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import lombok.Getter;
import lombok.Setter;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.io.ClassPathResource;
import org.springframework.web.client.RestTemplate;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.security.GeneralSecurityException;
import java.security.KeyFactory;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.cert.CertificateFactory;
import java.security.cert.X509Certificate;
import java.security.spec.PKCS8EncodedKeySpec;
import java.time.Duration;
import java.util.Base64;

@Configuration
@Getter
@Setter
public class Config {

    @Value("${base.url}")
    private String baseUrl;
    @Value("${private.key}")
    private String privateKey;
    @Value("${public.key}")
    private String publicKey;
    @Value("${merchant.code}")
    private String merchantCode;
    @Value("${client.code}")
    private String clientCode;
    @Value("${client.secret}")
    private String clientSecret;
    @Value("${gpay.public.key}")
    private String gpayPublicKey;


    @Bean
    public RestTemplate restTemplate(RestTemplateBuilder restTemplateBuilder) {
        return restTemplateBuilder
                .setConnectTimeout(Duration.ofSeconds(90))
                .setReadTimeout(Duration.ofSeconds(90))
                .build();
    }

    @Bean
    public Gson gsonSnackKey() {
        return new GsonBuilder()
                .setFieldNamingPolicy(FieldNamingPolicy.LOWER_CASE_WITH_UNDERSCORES)
                .create();
    }

    @Bean
    public PrivateKey privateKey() throws IOException, GeneralSecurityException {
        String pathFile = new ClassPathResource(privateKey).getFile().getPath();
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

    @Bean
    public String rawPublicKey() {
        try {
            String filePath = new ClassPathResource(publicKey).getFile().getAbsolutePath();
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

    @Bean
    public PublicKey publicKey() {
        try {
            String filePath = new ClassPathResource(publicKey).getFile().getAbsolutePath();
            Path pathCer = Paths.get(filePath);
            byte[] bytes = Files.readAllBytes(pathCer);
            ByteArrayInputStream byteArrayInputStream = new ByteArrayInputStream(bytes);

            CertificateFactory fact = CertificateFactory.getInstance("X.509");
            X509Certificate cer = (X509Certificate) fact.generateCertificate(byteArrayInputStream);
            return cer.getPublicKey();

        } catch (Exception var5) {
            var5.printStackTrace();
        }
        return null;
    }
}