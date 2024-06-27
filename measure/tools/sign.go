package tools

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
)

type Signer struct {
}

var signer *Signer = MakeSigner()

func GetSigner() *Signer {
	return signer
}

func MakeSigner() *Signer {
	return &Signer{}
}

func (s *Signer) Sign(fileHashs map[string][]byte, key []byte) map[string][][]byte {
	signResult := make(map[string][][]byte)
	for file, hash := range fileHashs {
		sign, err := RsaSign(key, hash)
		if err != nil {
			log.Fatalln("签名错误:", err)
			return nil
		}
		signResult[file] = [][]byte{hash, sign}
	}
	return signResult
}

func RsaSign(prvkey []byte, hash []byte) ([]byte, error) {
	block, _ := pem.Decode(prvkey)
	if block == nil {
		return nil, errors.New("decode private key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
}

func (s *Signer) Verify(key []byte, signResult map[string][][]byte) bool {
	for _, hashAndSig := range signResult {
		err := RsaVerifySign(key, hashAndSig[0], hashAndSig[1])
		if err != nil {
			log.Fatalln("验签错误:", err)
			return false
		}
	}

	return true
}

func (s *Signer) Verify2(key []byte, sigArray [][32]byte, hashArray [][32]byte) bool {
	for i, sig := range sigArray {
		tmp1 := sig[:]
		tmp2 := hashArray[i][:]
		err := RsaVerifySign(key, tmp2, tmp1)
		if err != nil {
			log.Fatalln("验签错误:", err)
			return false
		}
	}

	return true
}

func RsaVerifySign(pubkey []byte, hash, sig []byte) error {
	block, _ := pem.Decode(pubkey)
	if block == nil {
		return errors.New("decode public key error")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hash, sig)
}
