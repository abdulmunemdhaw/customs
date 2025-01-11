package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abdlmunemdhaw/customs/app"
	"github.com/abdlmunemdhaw/customs/config"
	"github.com/gin-gonic/gin"
)


func main() {
	config.Init();
	router := gin.Default();
	app.App(router);
	router.POST("/ping",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"msg":"PONG",
		})
	})
	log.Println("LISTEN ON PORT " + fmt.Sprintf("%v",config.ConfigData.Server.Port));
	err := router.Run(config.ConfigData.Server.Host + ":" + fmt.Sprintf("%v",config.ConfigData.Server.Port))
	if err != nil {
		log.Fatalln(err.Error());
	}
}