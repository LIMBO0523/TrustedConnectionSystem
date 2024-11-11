package measure

import (
	"TrustedConnectionSystem/contract"
	"TrustedConnectionSystem/fisco"
	"TrustedConnectionSystem/init"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"log"
)

type PinStorage struct {
	ContractAddress common.Address
	DeviceAddress   common.Address
	Session         *contract.PinStorageSession
	Client          *client.Client
}

var pinStorage *PinStorage

func GetPinStorage() *PinStorage {
	return pinStorage
}

func init() {
	conn := fisco.GetClient()
	session := GetPSInstance(conn)

	pinStorage = &PinStorage{
		ContractAddress: init.GetPinService(),
		DeviceAddress:   init.GetDeviceAddress(),
		Session:         session,
		Client:          conn,
	}
}

func GetPSInstance(conn *client.Client) *contract.PinStorageSession {
	// 调用合约IntegrityVerify.sol中的verifyAndStoreHashes方法
	address := init.GetPinService()
	instance, err := contract.NewPinStorage(address, conn)
	if err != nil {
		logrus.Errorln(" 获取智能合约实例PinStorage失败: ", err)
	}

	pinStorageSession := &contract.PinStorageSession{Contract: instance, CallOpts: *conn.GetCallOpts(), TransactOpts: *conn.GetTransactOpts()}
	return pinStorageSession
}

func (ps *PinStorage) CallGetPin() string {
	pin, err := ps.Session.GetPin()
	if err != nil {
		logrus.Errorln(" 获取Pin码失败: ", err)
	}
	return pin
}

func (ps *PinStorage) CallSetPin(newPin string) bool {
	_, _, err := ps.Session.SetPin(newPin)
	if err != nil {
		logrus.Errorln(" 设置Pin码失败: ", err)
	}
	return true
}

func (ps *PinStorage) DeployPinStorageContract() (bool, common.Address) {
	address, _, instance, err := contract.DeployPinStorage(ps.Client.GetTransactOpts(), ps.Client, "1234")
	if err != nil {
		log.Println("部署合约失败!")
		return false, address
	}
	session := &contract.PinStorageSession{Contract: instance, CallOpts: *ps.Client.GetCallOpts(), TransactOpts: *ps.Client.GetTransactOpts()}
	ps.Session = session
	ps.ContractAddress = address
	log.Println("部署合约成功！")
	return true, address
}
