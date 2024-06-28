package api

import (
	"TrustedConnectionSystem/conf"
	"TrustedConnectionSystem/measure"
	"TrustedConnectionSystem/measure/tools"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"log"
	"net/http"
)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func GenerateKey(c *gin.Context) {
	if tools.GetKeysPair() == nil {
		c.JSON(200, gin.H{
			"message": "fail",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

func SetFiles(c *gin.Context) {
	type RequestBody struct {
		Files   []string `json:"files" binding:"required"`
		Version string   `json:"version" binding:"required"`
	}
	var requestBody RequestBody

	// 绑定请求体到结构体
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tools.GetHasher().SetFiles(requestBody.Files, requestBody.Version)
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

func GetFiles(c *gin.Context) {
	hasher := tools.GetHasher()
	files, version, err := hasher.GetFiles()
	if err == nil {
		c.JSON(200, gin.H{
			"message": "success",
			"files":   files,
			"version": version,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "fail",
			"error":   err.Error(),
		})
	}
	return
}

func StartMeasure(c *gin.Context) {
	// 1. 度量文件获得度量值
	files := tools.GetHasher().HashFiles()
	hashAndSign := tools.GetSigner().Sign(files, tools.GetKeysPair().GetPrivateKey())
	log.Println("成功获得文件度量值！")

	// 2. 将度量值发送的智能合约进行验证
	success := measure.GetIntergrityVerify().CallVerifyAndStoreHashes(hashAndSign)
	log.Println("开始度量签名验证...")
	if !success {
		log.Println("度量签名验证失败")
		c.JSON(200, gin.H{
			"message": "度量签名验证失败",
		})
		return
	}

	// 3. 监听事件，获取验证结果
	result := measure.GetIntergrityVerify().WatchSetAddressToBool()
	if !result {
		log.Println("所有设备都同意接入")
	} else {
		log.Println("存在设备不同意接入")
	}

	// 4. 给对方返回验证结果

	// 如果没有监听到事件，说明验证通过
	if !result {
		c.JSON(200, gin.H{
			"message": "所有设备都同意接入",
			"address": conf.GetDeviceAddress(),
			"result":  true,
		})
		return
	}

	// 如果监听到事件，说明验证失败
	c.JSON(200, gin.H{
		"message": "存在设备不同意接入",
		"address": conf.GetDeviceAddress(),
		"result":  false,
	})
	return
}

func UploadPublicKey(c *gin.Context) {
	pair := tools.GetKeysPair()
	publicKey := pair.GetPublicKey()

	err := measure.GetIntergrityVerify().CallSetAddressToPublicKey(publicKey)
	if err != nil {
		log.Println("上传公钥失败")
		c.JSON(200, gin.H{
			"message": "上传公钥失败",
			"result":  false,
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "上传公钥成功",
			"result":  true,
		})
	}
}

func UploadFingerprint(c *gin.Context) {
	hasher := tools.GetHasher()
	hashFiles := hasher.HashFiles()
	files := make([]string, len(hashFiles))
	hashs := make([][32]byte, len(hashFiles))
	i := 0
	for file, bytes := range hashFiles {
		files[i] = file
		hashs[i] = [32]byte(bytes)
		i++
	}
	err := measure.GetIntergrityVerify().CallSetHash(files, hashs)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "上传指纹失败",
			"result":  false,
		})
		log.Fatalln("上传指纹失败")
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "上传指纹成功", "result": true})
		log.Println("上传指纹成功")
	}
}

func InitiateMeasurement(c *gin.Context) {
	type RequestPayload struct {
		IP   string `json:"IP" binding:"required"`
		PORT string `json:"Port" binding:"required"`
	}
	var request RequestPayload
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//url := "http://127.0.0.1:8080/startmeasure"
	url := "http://" + request.IP + ":" + request.PORT + "/startmeasure"

	// 发送请求
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 处理响应
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	type ResponseData struct {
		Message string         `json:"message"`
		Address common.Address `json:"address"`
		Result  bool           `json:"result"`
	}

	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func VerifyResult(c *gin.Context) {
	type RequestPayload struct {
		Address common.Address `json:"address" binding:"required"`
	}
	var paylaod RequestPayload
	if err := c.ShouldBindJSON(&paylaod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fromString, err := common.NewMixedcaseAddressFromString(addressString)
	//if err != nil {
	//	log.Fatalln("地址转换失败: ", err)
	//}
	//address := fromString.Address()

	ivc := measure.GetIntergrityVerify()
	isAddressTrue := ivc.CallIsAddressTrue(paylaod.Address)

	c.JSON(200, gin.H{
		"message": "验证结果",
		"address": paylaod.Address,
		"result":  isAddressTrue,
	})
}

func DeployContract(c *gin.Context) {
	ivc := measure.GetIntergrityVerify()
	success, address := ivc.DeployIntegrityVerify()

	if !success {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "部署失败",
			"result":  false,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "部署成功",
		"address": address,
		"result":  true,
	})
}
