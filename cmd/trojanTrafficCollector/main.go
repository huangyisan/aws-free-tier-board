package main

import (
	"time"
	"trojan-dashboard/app/trojanTraffic"
	myConfig "trojan-dashboard/pkg/configs"
)

func main() {
	myConfig.InitLog()
	for {
		trojanTraffic.RecordTrojanTraffic()
		<-time.Tick(time.Second * myConfig.TrojanTrafficCollectorConf.Duration)
	}

}
