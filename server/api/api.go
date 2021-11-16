package api

import (
	"fmt"
	"net/http"
	"server/conf"
	"server/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var err error

func getRouter() *gin.Engine {
	conf := conf.Get()
	api := gin.Default()
	api.Use(gin.Recovery())
	api.Use(middleware.CORS())
	g := api.Group(conf.Location)
	g.POST("/image", CreateImage)
	g.GET("/image/:id", GetImageByID)
	g.GET("/ping", ping)
	return api
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Start() {
	wait := 1
	conf := conf.Get()
	r := getRouter()
	log.Info().Msg(fmt.Sprintf("starting api in port %d", conf.Port))
	for {
		err = r.Run(fmt.Sprintf("%s:%d", conf.Ip, conf.Port))
		if wait == 32 {
			wait = 1
		} else {
			wait = wait * 2
		}
		log.Error().Err(err).Msg(fmt.Sprintf("cannot start rest api service, retrying in %d seconds", wait))
		time.Sleep(time.Duration(wait) * time.Second)
	}
}
