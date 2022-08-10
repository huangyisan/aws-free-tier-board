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
	LogLevel      string `yaml:"logLevel"`
	LogFile       string `yaml:"logFile"`
	ReportCaller  bool   `yaml:"reportCaller"`
	DisableColors bool   `yaml:"disableColors"`
	EnableFile    bool   `yaml:"enableFile"`
}

type ginConfig struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
}

var Aws []struct {
	Ak      string `yaml:"ak"`
	Sk      string `yaml:"sk"`
	Region  string `yaml:"region"`
	Account string `yaml:"account"`
}

type TrojanConfig []OneTrojanConfig

type OneTrojanConfig struct {
	Ip    string `yaml:"ip"`
	Port  string `yaml:"port"`
	Tag   string `yaml:"tag"`
	Group string `yaml:"group"`
}

type influxDBConfig struct {
	Url       string `yaml:"url"`
	AuthToken string `yaml:"authToken"`
	Org       string `yaml:"org"`
	Bucket    string `yaml:"bucket"`
}

type trojanTrafficCollectorConfig struct {
	Duration time.Duration `yaml:"duration"`
}
