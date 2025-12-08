package crypt

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
)

func EncryptRSA(datastring string) string {
	privkey := getPrivateKey()
	pemRes, _ := pem.Decode([]byte(privkey))
	privateKeyI, err := x509.ParsePKCS1PrivateKey(pemRes.Bytes)
	if err != nil {
		log.Errorf("Error x509.ParsePKCS1PrivateKey: %v\n", err)
		return ""
	}
	h := sha256.New()
	_, err = h.Write([]byte(datastring))
	if err != nil {
		return ""
	}

	sum := h.Sum(nil)

	sig, err := rsa.SignPKCS1v15(rand.Reader, privateKeyI, crypto.SHA256, sum)
	if err != nil {
		log.Errorf("Error rsa.SignPKCS1v15: %v\n", err)
		return ""
	}

	return base64.StdEncoding.EncodeToString(sig)
}

func VerifySHA256RSASign(message string, sig string) error {
	publicKeyB := getPublicKey()
	h := sha256.New()
	_, err := h.Write([]byte(message))
	if err != nil {
		return err
	}
	d := h.Sum(nil)

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyB)
	if err != nil {
		log.Errorf("Error parsing private key: %v\n", err)
		return err
	}

	sigDecode, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		log.Errorf("Error base64.StdEncoding.DecodeString: %v\n", err)
		return err
	}

	return rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, d, sigDecode)
}

// getPrivateKey
func getPrivateKey() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	pemFile, err := os.Open(dir + "/private_key.pem")
	defer pemFile.Close()
	if err != nil {
		return ""
	}

	pemFileInfo, err := pemFile.Stat()
	if err != nil {
		return ""
	}

	var size int64 = pemFileInfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(pemFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		return ""
	}

	return string(pembytes)
}

// GetCert
func GetCert() string {
	dir, _ := os.Getwd()
	b, err := os.ReadFile(dir + "/cert.crt") // just pass the file name
	fmt.Println(err)
	return string(b) // convert content to a 'string'
}

// getPublicKey
func getPublicKey() []byte {
	dir, err := os.Getwd()
	if err != nil {
		return nil
	}
	certPem, err := os.ReadFile(fmt.Sprintf("%v/%v", dir, "gpay_public_key.crt"))
	if err != nil {
		return nil
	}

	block, _ := pem.Decode(certPem)
	if block == nil {
		return nil
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil
	}

	// Convert public key to PKIX (SPKI) bytes
	pubBytes, err := x509.MarshalPKIXPublicKey(cert.PublicKey)
	if err != nil {
		return nil
	}

	return pubBytes
}
