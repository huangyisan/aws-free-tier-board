package main

import (
	"time"
	"trojan-dashboard/app/trojanTraffic"
	myConfig "trojan-dashboard/pkg/config"
)

func main() {
	myConfig.InitLog()
	for {
		<-time.Tick(time.Second * 300)
		trojanTraffic.RecordTrojanTraffic()
	}

}
