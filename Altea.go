package main

import (
	projectRouter "Altea/http/router"
	"Altea/services/FileService"
	"Altea/services/managers"
	"github.com/gin-gonic/gin"
)

func main() {

	//start gin server on port 9000
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	var fs FileService.FileServiceInterface = FileService.NewLocal("./data")

	fm := managers.NewFileManager(&fs)
	projectRouter.FileRouter(router, fm)
	if err := router.Run(":9000"); err != nil {
		panic(err)
	}

}
