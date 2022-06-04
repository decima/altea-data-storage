package main

import (
	projectRouter "Altea/http/router"
	"Altea/services/FileService"
	"Altea/services/managers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	//start gin server on port 9000
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	var fs FileService.FileServiceInterface = FileService.NewLocal("./data")

	fm := managers.NewFileManager(&fs)
	projectRouter.FileRouter(router, fm)
	host := "0.0.0.0:9000"
	log.Println("starting to listen to " + host)
	if err := router.Run(host); err != nil {
		panic(err)
	}

}
