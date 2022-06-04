package router

import (
	"Altea/http/controllers"
	"Altea/services/managers"
	"github.com/gin-gonic/gin"
)

func FileRouter(engine gin.IRouter, fileManager *managers.FileManager) {
	controller := controllers.NewFileController(fileManager)

	router := engine.Group("/files")

	router.GET("/*path", (*controller).Get)
	router.PUT("/*path", (*controller).Put)
	router.DELETE("/*path", (*controller).Delete)

}
