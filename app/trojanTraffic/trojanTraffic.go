package trojanTraffic

import (
	"sync"
	myConfig "trojan-dashboard/pkg/configs"
	"trojan-dashboard/pkg/logger"
	"trojan-dashboard/pkg/meInfluxdb"
	myTrojan "trojan-dashboard/pkg/meTrojan"
)

const (
	KB = 1 << (10 * (iota + 1))
	MB
	GB
)

func RecordTrojanTraffic() {
	var wg sync.WaitGroup
	trojanConfig := myConfig.TrojanConf
	influxDBConfig := myConfig.InfluxDBConf
	infdb := meInfluxdb.NewInfluxDBClient(influxDBConfig.Url, influxDBConfig.AuthToken, influxDBConfig.Org, influxDBConfig.Bucket)
	defer infdb.Close()
	for _, v := range *trojanConfig {
		wg.Add(1)
		go func(v myConfig.OneTrojanConfig) {
			defer wg.Done()
			tc := myTrojan.NewTrojanClient(v.Ip, v.Port)
			defer tc.Close()
			downloadTraffic, uploadTraffic := tc.ListAllTraffic()
			line := meInfluxdb.MakeLine("trojan,tag=%s,ip=%s download=%d,upload=%d", v.Tag, v.Ip, downloadTraffic/MB, uploadTraffic/MB)
			logger.Success.Msgf(line)
			infdb.WriteLineRecord(line)
		}(v)
	}
	wg.Wait()
}
