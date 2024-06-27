package main

import (
	"TrustedConnectionSystem/router"
)

func main() {
	r := router.ValidatorRouter()
	//err := r.Run(":" + conf.GetPort())
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
