package global

import "github.com/spf13/pflag"

var (
	CfgFile = pflag.StringP("config", "c", "", "set configuration file.")
)

func init() {
	pflag.Parse()
}
