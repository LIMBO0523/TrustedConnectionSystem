package measure

import (
	"TrustedConnectionSystem/conf"
	"TrustedConnectionSystem/contract"
	"TrustedConnectionSystem/fisco"
	"TrustedConnectionSystem/measure/tools"
	"context"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/client"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
	"time"
)

type IntergrityVerifyContract struct {
	ContractAddress common.Address
	DeviceAddress   common.Address
	Session         *contract.IntegrityVerifySession
	Client          *client.Client
}

var IntergrityVerify *IntergrityVerifyContract

func init() {
	conn := fisco.GetClient()
	session := GetContractInstance(conn)

	IntergrityVerify = &IntergrityVerifyContract{
		ContractAddress: conf.GetContractAddress(),
		DeviceAddress:   conf.GetDeviceAddress(),
		Session:         session,
		Client:          conn,
	}
}

func GetIntergrityVerify() *IntergrityVerifyContract {
	return IntergrityVerify
}

func GetContractInstance(client *client.Client) *contract.IntegrityVerifySession {
	// 调用合约IntegrityVerify.sol中的verifyAndStoreHashes方法
	address := conf.GetContractAddress()
	instance, err := contract.NewIntegrityVerify(address, client)
	if err != nil {
		logrus.Errorln(" 获取智能合约实例IntegrityVerify失败: ", err)
	}

	integrityVerifySession := &contract.IntegrityVerifySession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	return integrityVerifySession
}

// CallVerifyAndStoreHashes 调用智能合约验证文件完整性
func (ivc *IntergrityVerifyContract) CallVerifyAndStoreHashes(hashAndSign map[string][][]byte) bool {

	keys := make([]string, len(hashAndSign))
	hashs := make([][32]byte, len(hashAndSign))
	sigs := make([][32]byte, len(hashAndSign))
	i := 0
	for key, hs := range hashAndSign {
		keys[i] = key
		hashs[i] = [32]byte(hs[0])
		sigs[i] = [32]byte(hs[1])
		i++
	}

	check, _, _, err := ivc.Session.VerifyAndStoreHashes(keys, hashs, sigs)
	if err != nil {
		logrus.Errorln("合约VerifyAndStoreHashes调用失败: ", err)
	}
	return check
}

// WatchSetAddressToBool 监听事件measureFinished，2s后结束监听
func (ivc *IntergrityVerifyContract) WatchSetAddressToBool() bool {
	topic := common.BytesToHash(crypto.Keccak256([]byte("measureFinished()"))).Hex() // 事件的Topic计算，注意使用uint256及int256替代uint及int
	params := types.EventLogParams{
		FromBlock: -1,                                     // 最新的区块
		ToBlock:   -1,                                     // 最新的区块
		Addresses: []string{ivc.ContractAddress.String()}, // 要监听的合约地址
		Topics:    []string{topic},                        // 要监听的事件
	}

	isSomeDeviceDisagree := false
	// 监听事件
	_, err := ivc.Client.SubscribeEventLogs(context.Background(), params, func(i int, logs []types.Log) {
		isSomeDeviceDisagree = true
	})
	if err != nil {
		logrus.Errorln("合约监听事件失败: ", err)
	}

	time.Sleep(time.Second * 2)

	return isSomeDeviceDisagree
}

// CallIsAddressTrue 调用智能合约查看文件是否完整
func (ivc *IntergrityVerifyContract) CallIsAddressTrue(address common.Address) bool {
	isAddressTrue, err := ivc.Session.IsAddressTrue(address)
	if err != nil {
		logrus.Errorln("合约调用IsAddressTrue失败: ", err)
	}
	return isAddressTrue
}

// CallSetAddressToPublicKey 调用智能合约设置公钥
func (ivc *IntergrityVerifyContract) CallSetAddressToPublicKey(publicKey []byte) error {
	_, _, err := ivc.Session.SetAddressToPublicKey(ivc.DeviceAddress, publicKey)
	if err != nil {
		log.Println("合约调用SetAddressToPublicKey失败: ", err)
		return err
	}
	return nil
}

func (ivc *IntergrityVerifyContract) CallSetHash(filePath []string, hash [][32]byte) error {
	_, _, err := ivc.Session.SetMultipleHashes(filePath, hash)
	if err != nil {
		log.Println("合约调用SetHash失败: ", err)
		return err
	}
	return nil
}

func (ivc *IntergrityVerifyContract) CallGetHash(file string) ([32]byte, error) {
	hash, err := ivc.Session.GetHash(file)
	if err != nil {
		log.Println("合约调用GetHash失败: ", err)
		return [32]byte{}, err
	}
	return hash, nil
}

func (ivc *IntergrityVerifyContract) CallGetAddressToPublicKey() ([]byte, error) {
	key, err := ivc.Session.GetAddressToPublicKey(ivc.DeviceAddress)
	if err != nil {
		log.Println("合约调用GetAddressToPublicKey失败: ", err)
		return nil, err
	}
	return key, nil
}

func (ivc *IntergrityVerifyContract) CallSetAddressToBool(address common.Address) error {
	_, _, err := ivc.Session.SetAddressToBool(address, false)
	if err != nil {
		log.Println("合约调用SetAddressToBool失败: ", err)
		return err
	}
	return nil
}

func (ivc *IntergrityVerifyContract) WatchVerifyAndStoreHashes() {
	topic := common.BytesToHash(crypto.Keccak256([]byte("HashArrayStored(bytes32[], bytes, bytes32[], address)"))).Hex() // 事件的Topic计算，注意使用uint256及int256替代uint及int
	params := types.EventLogParams{
		FromBlock: -1,                                     // 最新的区块
		ToBlock:   -1,                                     // 最新的区块
		Addresses: []string{ivc.ContractAddress.String()}, // 要监听的合约地址
		Topics:    []string{topic},                        // 要监听的事件
	}

	type evidence struct {
		sigArray        [][32]byte
		senderPublicKey []byte
		hashArray       [][32]byte
		address         common.Address
	}
	// 监听事件
	channel := make(chan struct{})
	_, err := ivc.Client.SubscribeEventLogs(context.Background(), params, func(i int, logs []types.Log) {
		fmt.Println("监听到事件", logs)
		for _, v := range logs {
			var tmpABI abi.ABI
			var tmp evidence
			tmpABI, err := abi.JSON(strings.NewReader(contract.IntegrityVerifyBin))
			if err != nil {
				logrus.Errorln("abi解析失败: ", err.Error())
			}
			err = tmpABI.Unpack(&tmp, "HashArrayStored", v.Data)
			if err != nil {
				logrus.Errorln("数据解包失败: ", err.Error())
			}

			signer := tools.GetSigner()
			result := signer.Verify2(tmp.senderPublicKey, tmp.sigArray, tmp.hashArray)

			if !result {
				err := ivc.CallSetAddressToBool(tmp.address)
				if err != nil {
					logrus.Errorln("合约调用SetAddressToBool失败: ", err)
				}
			}
		}
	})
	if err != nil {
		logrus.Errorln("合约监听事件失败: ", err)
	}

	<-channel
}

func (ivc *IntergrityVerifyContract) DeployIntegrityVerify() (bool, common.Address) {
	address, _, instance, err := contract.DeployIntegrityVerify(ivc.Client.GetTransactOpts(), ivc.Client)
	if err != nil {
		log.Println("部署合约失败!")
		return false, address
	}
	session := &contract.IntegrityVerifySession{Contract: instance, CallOpts: *ivc.Client.GetCallOpts(), TransactOpts: *ivc.Client.GetTransactOpts()}
	ivc.Session = session
	ivc.ContractAddress = address
	log.Println("部署合约成功！")
	return true, address
}
