package config

var (
	LogrusConf *LogrusConfig
	GinConf    *GinConfig
)

type LogrusConfig struct {
	LogLevel      string
	LogFile       string
	ReportCaller  bool
	DisableColors bool
	EnableFile    bool
}

type GinConfig struct {
	Ip   string
	Port string
	Mode string
}

var Aws []struct {
	Ak string
	Sk string
}

type TrojanConfig struct {
	Ip       string
	Port     string
	Tag      string
	Password string
}
