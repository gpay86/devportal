package com.gpay.gpaysample.GpaySampleApplication.utils;

import java.io.ByteArrayInputStream;
import java.io.DataInputStream;
import java.io.File;
import java.io.FileInputStream;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.security.KeyFactory;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.Signature;
import java.security.cert.CertificateFactory;
import java.security.cert.X509Certificate;
import java.security.interfaces.RSAPublicKey;
import java.security.spec.PKCS8EncodedKeySpec;
import java.util.Base64;

public class CryptoUtils {
    private static final String SIGNATURE_ALGORITHM = "SHA256withRSA";
    private static final String ENCODING = "UTF-8";

    private CryptoUtils() {
    }

    public static String createSignature(String path, String data) throws Exception {
        PrivateKey privateKey = getPemPrivateKey(path);
        Signature privateSignature = Signature.getInstance(SIGNATURE_ALGORITHM);
        privateSignature.initSign(privateKey);
        privateSignature.update(data.getBytes(ENCODING));

        byte[] signature = privateSignature.sign();

        return Base64.getEncoder().encodeToString(signature);
    }

    public static boolean verifySignature(String pathKey, String data, String sign) throws Exception {

        RSAPublicKey publicKey = (RSAPublicKey) getPemPublicKey(pathKey);
        Signature signatureAlgorithm = Signature.getInstance(SIGNATURE_ALGORITHM);
        signatureAlgorithm.initVerify(publicKey);
        signatureAlgorithm.update(data.getBytes(StandardCharsets.UTF_8));
        byte[] signatureBytes = Base64.getDecoder().decode(sign);
        return signatureAlgorithm.verify(signatureBytes);
    }


    public static PrivateKey getPemPrivateKey(String filename) throws Exception {
        File f = new File(filename);
        FileInputStream fis = new FileInputStream(f);
        DataInputStream dis = new DataInputStream(fis);
        byte[] keyBytes = new byte[(int) f.length()];
        dis.readFully(keyBytes);
        dis.close();

        String temp = new String(keyBytes);
        String privKeyPEM = temp.replace("-----BEGIN PRIVATE KEY-----", "")
                .replace("-----END PRIVATE KEY-----", "")
                .replace("\n", "").replace("\r", "");

        PKCS8EncodedKeySpec spec = new PKCS8EncodedKeySpec(Base64.getDecoder().decode(privKeyPEM.getBytes()));
        KeyFactory kf = KeyFactory.getInstance("RSA");
        return kf.generatePrivate(spec);
    }

    public static PublicKey getPemPublicKey(String filename) throws Exception {
        try {
            Path pathCer = Paths.get(filename);
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
