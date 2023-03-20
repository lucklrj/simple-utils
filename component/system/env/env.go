package env

import (
	"flag"
	"os"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Envs map[string]string
var envPath *string

func init() {
	envPath = flag.String("env_path", "./.env", "配置文件地址")
	flag.Parse()
}
func ParseEnv() {

	color.Green("开始解析.env配置文件")

	viper.SetConfigFile(*envPath)
	err := viper.ReadInConfig()
	if err != nil {
		color.Red("解析.env文件错误：" + err.Error())
		os.Exit(0)
	}
	load()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		load()
	})

}
func load() {
	var result map[string]string
	viper.Unmarshal(&result)
	Envs = result
}
