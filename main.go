package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"trojan-dashboard/config"
	"trojan-dashboard/http/router"
)

func main() {
	// init logger
	config.InitLog()
	aws := config.Aws
	for _, v := range aws {
		fmt.Println(v)
	}
	gin.SetMode(config.GinConf.Mode)
	g := gin.New()

	router.Load(g)
	log.Info(http.ListenAndServe(fmt.Sprintf("%s:%s", config.GinConf.Ip, config.GinConf.Port), g))

}
