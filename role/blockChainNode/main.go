package main

import (
	"TrustedConnectionSystem/measure"
	"fmt"
)

func main() {
	ivc := measure.GetIntergrityVerify()
	fmt.Println("开始监听事件HashArrayStored")
	ivc.WatchVerifyAndStoreHashes()
	fmt.Println("事件HashArrayStored监听结束")
}
