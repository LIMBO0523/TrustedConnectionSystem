package main

import (
	"TrustedConnectionSystem/measure"
	"TrustedConnectionSystem/router"
	"fmt"
	"log"
)

func WatchVerifyAndStoreHashes() {
	ivc := measure.GetIntergrityVerify()
	fmt.Println("开始监听事件HashArrayStored")
	ivc.WatchVerifyAndStoreHashes()
	fmt.Println("事件HashArrayStored监听结束")
}
func main() {
	go WatchVerifyAndStoreHashes()
	r := router.BlockChainNodeRouter()
	err := r.Run(":8082")
	if err != nil {
		log.Fatalln(err.Error())
	}
}
