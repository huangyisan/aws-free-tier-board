package main

import (
	"time"
	"trojan-dashboard/app/trojanTraffic"
	myConfig "trojan-dashboard/pkg/configs"
	"trojan-dashboard/pkg/logger"
)

func main() {
	logger.InitLog()
	for {
		trojanTraffic.RecordTrojanTraffic()
		<-time.Tick(time.Second * myConfig.TrojanTrafficCollectorConf.Duration)
	}

}
