package fisco

import (
	"TrustedConnectionSystem/init"
	"context"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/sirupsen/logrus"
	"log"
)

var conn *client.Client

func init() {
	c, err := Conn(init.GetPEMFILE())
	if err != nil {
		logrus.Errorln(" 连接FISCO BCOS失败: ", err)
	}
	conn = c
}

func GetClient() *client.Client {
	return conn
}

func Conn(PEMFile string) (*client.Client, error) {
	privateKey, _, err := client.LoadECPrivateKeyFromPEM(PEMFile)
	if err != nil {
		log.Fatalln("get privateKey fail")
		return nil, err
	}
	// 连接方法一，输入参数
	//config.json := &client.Config{IsSMCrypto: false, GroupID: "group0",
	//	PrivateKey: privateKey, Host: "127.0.0.1", Port: 20200, TLSCaFile: "./conf/ca.crt", TLSKeyFile: "./conf/sdk.key", TLSCertFile: "./conf/sdk.crt"}
	//client, err := client.DialContext(context.Background(), config.json)
	// 连接方法二、配置文件
	//cli, err := client.Dial("/home/limbo/TrustedConnectionSystem/fisco/config", "group0", privateKey)
	cli, err := client.Dial("./fisco/config", "group0", privateKey)

	if err != nil {
		log.Fatalln("get client fail")
		return nil, err
	}

	return cli, nil
}

func GetBlockNumber(c *client.Client) (int64, error) {
	number, err := c.GetBlockNumber(context.Background())
	if err != nil {
		log.Fatalln("get blockNumber fail")
		return -1, err
	} else {
		log.Println("get blockNumber: ", number)
	}
	return number, nil
}
