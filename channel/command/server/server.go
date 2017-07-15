package server

import (
	"fmt"
	"log"

	channelConfig "github.com/mkfsn/shyuukan-program/channel/config"
	"github.com/mkfsn/shyuukan-program/channel/controller"
	"github.com/mkfsn/shyuukan-program/channel/security"

	"github.com/gin-gonic/gin"
	"github.com/mkfsn/shyuukan-program/channel/model"
)

func Run(host string, port int, configFile string) {
	var rest controller.Controller

	config, err := channelConfig.ReadConfig(configFile)
	if err != nil {
		config = channelConfig.DefaultConfig
		log.Println("Cannot read config:", err)
		log.Println("Use default config:", config)
	}

	db := model.OpenDB(config.Server.Database, config.Server.Debug, config.Channels)
	router := gin.New()

	router.Use(security.Middleware(config))

	rest = controller.NewChannelController(db)
	rest.AddRoutes("/api/channels", router)

	rest = controller.NewProgramController(db)
	rest.AddRoutes("/api/programs", router)

	router.Run(fmt.Sprintf("%s:%d", host, port))
}