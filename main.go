package main

import (
	"ytt-societyservice/config"
	"ytt-societyservice/routers"
)

func main() {
	config.Init()
	router := routers.InitRouter()
	router.Run(":8083")

}
