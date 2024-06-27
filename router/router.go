package router

import (
	"TrustedConnectionSystem/api"
	"github.com/gin-gonic/gin"
)

func ApplicantRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", api.Pong)

	// 上传公钥
	r.PUT("/uploadpublickey", api.UploadPublicKey)
	// 上传指纹(会把当前的指定文件的hash值作为指纹上传到智能合约)
	r.PUT("/uploadfingerprint", api.UploadFingerprint)
	// 开始度量
	r.POST("/startmeasure", api.StartMeasure)

	return r
}

func ValidatorRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", api.Pong)

	// 发起度量
	r.POST("/initiatemeasurement", api.InitiateMeasurement)

	// 查看验证结果
	r.POST("/verifyresult", api.VerifyResult)

	return r
}

func BlockChainNodeRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", api.Pong)

	//	部署智能合约
	r.GET("/deploycontract", api.DeployContract)

	return r
}
