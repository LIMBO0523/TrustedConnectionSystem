package tools

import (
	"TrustedConnectionSystem/conf"
	"crypto/sha256"
	"io"
	"log"
	"os"
	"strings"
)

type Hasher struct {
	files         map[string]struct{}
	systemVersion string
}

var hasher *Hasher

func init() {
	hasher = MakeHasher()
	hasher.SetFiles(conf.GetFilesAndVersion())
}

func GetHasher() *Hasher {
	return hasher
}

// MakeHasher 创建一个Hasher
func MakeHasher() *Hasher {
	return &Hasher{files: make(map[string]struct{})}
}

// SetFiles 设置要进行哈希的文件名
func (h *Hasher) SetFiles(files []string, version string) {
	for _, file := range files {
		h.files[file] = struct{}{}
	}
	h.systemVersion = version
}

// GetFiles 获取要进行哈希的文件名
func (h *Hasher) GetFiles() (files []string, version string, err error) {
	files = make([]string, 0)
	for file, _ := range h.files {
		files = append(files, file)
	}

	return files, h.systemVersion, nil
}

// HashFiles 对指定的文件进行sha256哈希
func (h *Hasher) HashFiles() map[string][]byte {
	result := make(map[string][]byte)

	for filePath, _ := range h.files {
		fileHash := h.HashFile(filePath)
		split := strings.Split(filePath, "/")
		filePath = split[len(split)-1] + "-" + h.systemVersion
		result[filePath] = fileHash
	}

	return result
}

func (h *Hasher) HashFile(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("无法打开文件:", err)
		return nil
	}
	defer file.Close()

	hash := sha256.New()

	buffer := make([]byte, 2048)

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Println("文件" + filePath + "读取完毕")
				break
			}
			log.Fatalln("文件读取出错:", err)
			return nil
		}
		hash.Write(buffer[:n])
	}

	return hash.Sum(nil)
}
