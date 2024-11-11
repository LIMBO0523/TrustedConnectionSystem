package init

import (
	"bufio"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	data, err := os.ReadFile("/home/limbo/TrustedConnectionSystem/conf/config.json")
	if err != nil {
		return
	}

	var info Config
	err = json.Unmarshal([]byte(data), &info)
	if err != nil {
		logrus.Errorln(" 解析配置文件失败: ", err)
	}
	fmt.Println(info.PemFile)
	fmt.Println(info.ContractAddress)
	fmt.Println(info.Files)
	fmt.Println(info.Version)
}

func TestReadLine(t *testing.T) {
	file, err := os.OpenFile("./config.json", os.O_RDONLY, 0666)
	if err != nil {
		logrus.Errorln("打开config文件失败:", err)
	}
	reader := bufio.NewReader(file)
	line, _, err := reader.ReadLine()
	line, _, err = reader.ReadLine()
	line, _, err = reader.ReadLine()

	fmt.Println(line)
}
