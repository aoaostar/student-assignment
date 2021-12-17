package bootstrap

import (
	"StudentAssignment/config"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() {
	readConfig("config.yaml", &config.Conf)
}
func readConfig(filepath string, Conf interface{}) {
	vipConfig := viper.New()

	vipConfig.SetConfigFile(filepath) // 指定配置文件路径
	// 查找并读取配置文件
	if err := vipConfig.ReadInConfig(); err != nil {
		log.Fatalf("读取配置出错: %s \n", err)
	}
	if err := vipConfig.Unmarshal(Conf); err != nil {
		log.Fatalf("解析配置出错: %s \n", err)
	}
	vipConfig.WatchConfig()
	vipConfig.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		log.Debug("config file changed:", e.Name)
		readConfig(filepath, Conf)
	})
}
