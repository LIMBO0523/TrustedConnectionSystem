package tools

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"path/filepath"
)

type KeyPair struct {
	privateKey []byte
	publicKey  []byte
}

var keyPair *KeyPair

func init() {
	keyPair = GetOrGenerateKeys()
}

func (pair *KeyPair) GetPrivateKey() []byte {
	return keyPair.privateKey
}

func (pair *KeyPair) GetPublicKey() []byte {
	return pair.publicKey
}

func GetKeysPair() *KeyPair {
	return keyPair
}

func generateKeys() *KeyPair {
	//rsa 密钥文件产生
	prvKey, pubKey := GenRsaKey()
	myKeyPair := &KeyPair{privateKey: prvKey, publicKey: pubKey}
	err := myKeyPair.SavePrivateKey()
	if err != nil {
		log.Fatalln("无法保存私钥:", err)
	}
	err = myKeyPair.SavePublicKey()
	if err != nil {
		log.Println("无法保存公钥:", err)
	}
	return myKeyPair
}

func readPublicKey() ([]byte, error) {
	filePath := "./myKeys/publicKey"
	//filePath := "D:\\Work\\openHarmony\\privateKey"
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func readPrivateKey() ([]byte, error) {
	filePath := "./myKeys/privateKey"
	//filePath := "D:\\Work\\openHarmony\\privateKey"
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetOrGenerateKeys 从本地获取私钥，或者生成公私钥
func GetOrGenerateKeys() *KeyPair {
	// 读取已有的公钥私钥
	privateKey, err := readPrivateKey()
	publicKey, _ := readPublicKey()
	if err != nil {
		log.Println("无法读取私钥:", err)
		log.Println("正在生成公私钥...")
		// 生成公钥私钥
		return generateKeys()
	} else {
		return &KeyPair{privateKey: privateKey, publicKey: publicKey}
	}
}

// SavePrivateKey 将私钥保存至本地
func (pair *KeyPair) SavePrivateKey() error {
	dir := "./myKeys"
	//dir := "D:\\Work\\openHarmony"
	filename := "privateKey"
	filePath := filepath.Join(dir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalln("无法创建文件privateKey:", err)
		return err
	}

	defer file.Close()

	_, err = file.Write(pair.privateKey)
	if err != nil {
		log.Fatalln("无法写入文件privateKey:", err)
		return err
	}

	log.Println("privateKey文件保存成功")
	return nil
}

func (pair *KeyPair) SavePublicKey() error {
	dir := "./myKeys"
	//dir := "D:\\Work\\openHarmony"
	filename := "publicKey"
	filePath := filepath.Join(dir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalln("无法创建文件publicKey:", err)
		return err
	}

	defer file.Close()

	_, err = file.Write(pair.publicKey)
	if err != nil {
		log.Fatalln("无法写入文件publicKey:", err)
		return err
	}

	log.Println("publicKey文件保存成功")
	return nil
}

// GenRsaKey RSA公钥私钥产生
func GenRsaKey() (prvkey, pubkey []byte) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey = pem.EncodeToMemory(block)
	return
}
