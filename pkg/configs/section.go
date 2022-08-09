package configs

import "time"

var (
	LogrusConf                 *logrusConfig
	GinConf                    *ginConfig
	TrojanConf                 *TrojanConfig
	InfluxDBConf               *influxDBConfig
	TrojanTrafficCollectorConf *trojanTrafficCollectorConfig
)

type logrusConfig struct {
	LogLevel      string
	LogFile       string
	ReportCaller  bool
	DisableColors bool
	EnableFile    bool
}

type ginConfig struct {
	Ip   string
	Port string
	Mode string
}

var Aws []struct {
	Ak string
	Sk string
}

type TrojanConfig []OneTrojanConfig

type OneTrojanConfig struct {
	Ip       string
	Port     string
	Tag      string
	Password string
}

type influxDBConfig struct {
	Url       string
	AuthToken string
	Org       string
	Bucket    string
}

type trojanTrafficCollectorConfig struct {
	Duration time.Duration
}
