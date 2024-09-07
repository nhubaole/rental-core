package initialize

import (
	"smart-rental/global"

	"github.com/spf13/viper"
)


func LoadConfig() (err error) {

	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&global.Config)
	return
}

