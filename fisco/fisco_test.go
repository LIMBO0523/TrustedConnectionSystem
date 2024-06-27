package fisco

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestConn(t *testing.T) {
	client, err := Conn("/home/limbo/fisco/accounts/0x53eab434a25cfb5949c6666950d0a374976d2c41.pem")
	if err != nil {
		logrus.Errorln(err)
	}

	number, err := GetBlockNumber(client)
	if err != nil {
		logrus.Errorln(err)
	}
	logrus.Infoln(number)
}
