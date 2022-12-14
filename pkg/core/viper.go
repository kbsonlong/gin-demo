package core

import (
	"fmt"

	"github.com/kbsonlong/gin-demo/pkg/global"
	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		config = "config.yaml"
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	//viper.AddConfigPath(".")
	v.SetConfigFile(config)
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	// 实时读取配置文件
	v.OnConfigChange(func(e fsnotify.Event) {
		global.LOG.Info("Config file changed:", zap.String("filename", e.Name))
	})
	// 获取环境变量
	v.AutomaticEnv()

	// 把配置解析到struct
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
