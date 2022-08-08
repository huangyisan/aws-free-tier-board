package main

import (
	"time"
	"trojan-dashboard/app/trojanTraffic"
	myConfig "trojan-dashboard/pkg/configs"
)

func main() {
	myConfig.InitLog()
	for {
		<-time.Tick(time.Second * 300)
		trojanTraffic.RecordTrojanTraffic()
	}

}
