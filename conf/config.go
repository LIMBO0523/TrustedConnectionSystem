package conf

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Port            string   `json:"port"`
	PemFile         string   `json:"pem_file"`
	ContractAddress string   `json:"contract_address"`
	DeviceAddress   string   `json:"device_address"`
	Files           []string `json:"files"`
	Version         string   `json:"version"`
	PinService      string   `json:"pin_service"`
}

var config = new(Config)

func init() {
	//data, err := os.ReadFile("/home/limbo/TrustedConnectionSystem/conf/config.json")
	data, err := os.ReadFile("./conf/config.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		logrus.Errorln(" 解析配置文件失败: ", err)
	}
}

func GetPEMFILE() string {
	return config.PemFile
}

func GetContractAddress() common.Address {
	fromString, err := common.NewMixedcaseAddressFromString(config.ContractAddress)
	if err != nil {
		logrus.Errorln(" 合约地址转换失败: ", err)
	}

	address := fromString.Address()
	return address
}

func GetPinService() common.Address {
	fromString, err := common.NewMixedcaseAddressFromString(config.PinService)
	if err != nil {
		logrus.Errorln(" 合约地址转换失败: ", err)
	}

	address := fromString.Address()
	return address
}

func GetDeviceAddress() common.Address {
	fromString, err := common.NewMixedcaseAddressFromString(config.DeviceAddress)
	if err != nil {
		logrus.Errorln(" 账户地址转换失败: ", err)
	}

	address := fromString.Address()
	return address
}

func GetFilesAndVersion() ([]string, string) {
	return config.Files, config.Version
}
