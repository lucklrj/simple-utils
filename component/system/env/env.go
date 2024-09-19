package env

import (
	"errors"
	"flag"

	"github.com/fsnotify/fsnotify"
	systemHelper "github.com/lucklrj/simple-utils/helper/system"
	"github.com/spf13/viper"
)

var Envs map[string]string
var envPath *string

func init() {
	envPath = flag.String("env_path", "./.env", "配置文件地址")
}
func Load(callback func(map[string]string)) error {
	if systemHelper.IsFileExist(*envPath) == false {
		return errors.New(".env文件不存在")
	}
	viper.SetConfigFile(*envPath)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	load(callback)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		load(callback)
	})
	return nil
}
func load(callback func(map[string]string)) {
	var result map[string]string
	viper.Unmarshal(&result)
	callback(result)
}
