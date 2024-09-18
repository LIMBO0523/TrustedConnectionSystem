package measure

import (
	"TrustedConnectionSystem/measure/tools"
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestIntergrityVerifyContract_CallSetAddressToPublicKey(t *testing.T) {

	ivc := GetIntergrityVerify()

	err := ivc.CallSetAddressToPublicKey(tools.GetKeysPair().GetPublicKey())
	if err != nil {
		return
	}
	key, _ := ivc.Session.GetAddressToPublicKey(ivc.DeviceAddress)

	fmt.Println(string(key))
	fmt.Println(string(tools.GetKeysPair().GetPublicKey()))
	if bytes.Equal(key, tools.GetKeysPair().GetPublicKey()) {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}

func TestIntergrityVerifyContract_CallSetHash(t *testing.T) {
	ivc := GetIntergrityVerify()

	hashFiles := tools.GetHasher().HashFiles()
	filePath := make([]string, len(hashFiles))
	hash := make([][32]byte, len(hashFiles))
	i := 0
	for key, value := range hashFiles {
		filePath[i] = key
		hash[i] = [32]byte(value)
		fmt.Println(filePath[i], hash[i])
	}

	err := ivc.CallSetHash(filePath, hash)
	if err != nil {
		log.Println(err)
	}

	re, _ := ivc.Session.GetHash("IntegrityVerify.go-1.0.0")

	fmt.Println(re)
}

func TestIntergrityVerifyContract_CallGetHash(t *testing.T) {
	ivc := GetIntergrityVerify()

	hash, err := ivc.CallGetHash("IntegrityVerify.go-1.0.0")
	if err != nil {
		log.Println(err)
	}
	filehash := tools.GetHasher().HashFile("/home/limbo/TrustedConnectionSystem/contract/IntegrityVerify.go")
	if hash == [32]byte(filehash) {
		fmt.Println("success1")
	}

	hash, err = ivc.CallGetHash("IntegrityVerify.sol-1.0.0")
	if err != nil {
		log.Println(err)
	}
	filehash = tools.GetHasher().HashFile("/home/limbo/TrustedConnectionSystem/contract/IntegrityVerify.sol")
	if hash == [32]byte(filehash) {
		fmt.Println("success2")
	}
}

func TestIntergrityVerifyContract_CallGetAddressToPublicKey(t *testing.T) {
	ivc := GetIntergrityVerify()
	key, err := ivc.CallGetAddressToPublicKey()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%c", key)
}

func TestIntergrityVerifyContract_CallVerifyAndStoreHashes(t *testing.T) {
	ivc := GetIntergrityVerify()
	hashFiles := tools.GetHasher().HashFiles()
	sign := tools.GetSigner().Sign(hashFiles, tools.GetKeysPair().GetPrivateKey())

	result := ivc.CallVerifyAndStoreHashes(sign)

	fmt.Println(result)
}
