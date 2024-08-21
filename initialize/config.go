package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/hespecial/gin-mall/config"
	"github.com/hespecial/gin-mall/global"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func LoadConfig() (conf *config.Config) {
	conf = new(config.Config)
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error viper ReadInConfig: %s", err))
	}
	err = viper.Unmarshal(conf)
	if err != nil {
		panic(fmt.Sprintf("Fatal error viper Unmarshal: %s", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		global.Log.Warn("Config file changed:", zap.String("config", e.Name))
	})
	viper.WatchConfig()
	return
}
