package main

import (
	"TrustedConnectionSystem/router"
)

func main() {
	/*
		http请求：
		1. 能够生成密钥，将私钥存储在可信环境中，将公钥发布到区块链上
		2. 能够指定要度量的文件
		3. 能够发起度量
			1. 度量文件获得度量值
			2. 将度量值发送的智能合约进行验证
			3. 监听事件，获取验证结果
		4. 监听事件，对度量签名进行验签，并将验签结果发送到区块链上
	*/
	r := router.ApplicantRouter()
	//err := router.Run(":" + conf.GetPort())
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
