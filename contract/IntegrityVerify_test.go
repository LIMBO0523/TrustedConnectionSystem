package contract

import (
	"TrustedConnectionSystem/conf"
	"TrustedConnectionSystem/fisco"
	"TrustedConnectionSystem/measure/tools"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestIntegrityVerify(t *testing.T) {
	conn, err := fisco.Conn(conf.GetPEMFILE())
	if err != nil {
		logrus.Errorln(" 连接FISCO BCOS失败: ", err)
	}

	mixedcaseAddressFromString, err := common.NewMixedcaseAddressFromString("0x21f3e1a1b648834aac88e0a2b9a78315db9ab8af")
	if err != nil {
		logrus.Errorln(" 0x21f3e1a1b648834aac88e0a2b9a78315db9ab8af地址转换失败: ", err)
	}
	address := mixedcaseAddressFromString.Address()
	instance, err := NewIntegrityVerify(address, conn)
	if err != nil {
		logrus.Errorln(" 获取智能合约实例IntegrityVerify失败: ", err)
	}

	intergrityVerify := &IntegrityVerifySession{Contract: instance, CallOpts: *conn.GetCallOpts(), TransactOpts: *conn.GetTransactOpts()}
	hasher := tools.GetHasher()
	bytes := hasher.HashFile("./IntegrityVerify.sol")
	_, _, err = intergrityVerify.SetHash("intergrityVerify", [32]byte(bytes))
	if err != nil {
		logrus.Errorln(" 调用智能合约SetHash失败: ", err)
	}

	hash, err := intergrityVerify.GetHash("intergrityVerify")
	if err != nil {
		logrus.Errorln(" 调用智能合约GetHash失败: ", err)
	}
	fmt.Println("hash值为: ", hash)
}
