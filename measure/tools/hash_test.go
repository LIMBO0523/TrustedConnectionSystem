package tools

import (
	"fmt"
	"testing"
)

func TestHasher_SetFiles(t *testing.T) {
	h := MakeHasher()
	files := []string{"D:\\Work\\openHarmony\\privateKey", "D:\\Work\\openHarmony\\完整性度量.md"}
	h.SetFiles(files, "3.2")
	fmt.Println(h.files)
}

func TestHasher_HashFile(t *testing.T) {
	h := MakeHasher()
	filePath := "D:\\Work\\openHarmony\\privateKey"

	hash := h.HashFile(filePath)
	fmt.Printf("%x", hash)
}

func TestHasher_HashFiles(t *testing.T) {
	h := MakeHasher()
	files := []string{"D:\\Work\\openHarmony\\privateKey", "D:\\Work\\openHarmony\\完整性度量.md"}
	h.SetFiles(files, "3.2")

	hashFiles := h.HashFiles()
	fmt.Println(hashFiles)
}

func TestHasher(t *testing.T) {
	getHasher := GetHasher()
	fmt.Println(getHasher.files)
	fmt.Println(getHasher.systemVersion)
}
