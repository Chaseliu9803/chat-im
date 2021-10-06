package config

import "github.com/spf13/viper"

var (
	CnfFile = "conf/config.yaml"
)

func SetConfig()  {
	viper.SetConfigFile(CnfFile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
