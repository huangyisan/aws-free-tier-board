package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"trojan-dashboard/http/router"
	config2 "trojan-dashboard/pkg/configs"
)

func main() {
	// init logger
	config2.InitLog()
	aws := config2.Aws
	for _, v := range aws {
		fmt.Println(v)
	}
	gin.SetMode(config2.GinConf.Mode)
	g := gin.New()

	router.Load(g)
	log.Info(http.ListenAndServe(fmt.Sprintf("%s:%s", config2.GinConf.Ip, config2.GinConf.Port), g))

}
