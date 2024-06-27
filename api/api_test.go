package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestSetFiles(t *testing.T) {
	type RequestBody struct {
		Files   []string `json:"files"`
		Version string   `json:"version"`
	}

	url := "http://172.24.185.80:8080/setfiles"

	// 创建请求数据
	data := RequestBody{
		Files:   []string{"file1.txt", "file2.txt", "file3.txt"},
		Version: "2.0",
	}

	// 将请求数据编码为JSON格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// 创建一个新的POST请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 处理响应
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Request successful")
	} else {
		fmt.Println("Request failed with status code:", resp.StatusCode)
	}
}

func TestStartMeasure(t *testing.T) {
	StartMeasure(&gin.Context{})
	//url := "http://172.24.185.80:8080/startmeasure"
	//req, err := http.NewRequest("POST", url, nil)
	//if err != nil {
	//	fmt.Println("Error creating request:", err)
	//	return
	//}
	//req.Header.Set("Content-Type", "application/json")
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println("Error sending request:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//// 处理响应
	//if resp.StatusCode == http.StatusOK {
	//	fmt.Println("Request successful")
	//	fmt.Println(resp.Body)
	//} else {
	//	fmt.Println("Request failed with status code:", resp.StatusCode)
	//	fmt.Println(resp.Body)
	//}
}

func TestInitiateMeasurement(t *testing.T) {
	//url := "http://172.24.185.80:8080/initiatemeasurement"
	//
	//type requestPayload struct {
	//	IP   string `json:"IP"`
	//	Port string `json:"Port"`
	//}
	//payload := requestPayload{
	//	IP:   "172.24.185.80",
	//	Port: "8080",
	//}
	//
	//jsonData, err := json.Marshal(payload)
	//if err != nil {
	//	log.Fatalf("Error marshalling JSON: %v", err)
	//}
	//
	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	//if err != nil {
	//	fmt.Println("Error creating request:", err)
	//	return
	//}
	//req.Header.Set("Content-Type", "application/json")
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println("Error sending request:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//// 处理响应
	//if resp.StatusCode == http.StatusOK {
	//	fmt.Println("Request successful")
	//	fmt.Println(resp.Body)
	//} else {
	//	fmt.Println("Request failed with status code:", resp.StatusCode)
	//	fmt.Println(resp.Body)
	//}
}
