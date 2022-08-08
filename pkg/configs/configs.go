package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"trojan-dashboard/global"
)

type Setting struct {
	vp *viper.Viper
}

func SetupConfig(cfgFile string) (*Setting, error) {
	vp := viper.New()
	if cfgFile == "" {
		vp.AddConfigPath("conf")
		vp.SetConfigName("app")
	} else {
		vp.SetConfigFile(cfgFile)
	}
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func init() {
	if err := initConfig(); err != nil {
		Failed.Msgf(fmt.Sprintf("init app config failed %v", err))
	}
}

func initConfig() error {
	setting, err := SetupConfig(*global.CfgFile)
	if err != nil {
		return err
	}

	err = setting.SetConf("logrus", &LogrusConf)
	if err != nil {
		return err
	}

	err = setting.SetConf("gin", &GinConf)
	if err != nil {
		return err
	}

	err = setting.SetConf("aws", &Aws)
	if err != nil {
		return err
	}

	err = setting.SetConf("trojan", &TrojanConf)
	if err != nil {
		return err
	}

	err = setting.SetConf("influxDB", &InfluxDBConf)
	if err != nil {
		return err
	}

	return nil
}

func (s *Setting) SetConf(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
