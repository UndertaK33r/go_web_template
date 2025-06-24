package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 全局变量，用来保存程序的所有配置信息
// 减少viper的使用次数，直接定义一个全局变量Conf，用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	// App配置信息
	*AppInfo `mapstructure:"app"`
	// 日志配置信息
	*LogConfig `mapstructure:"log"`
	// MySQL配置信息
	*MySQLConfig `mapstructure:"mysql"`
	// Redis配置信息
	*RedisConfig `mapstructure:"redis"`
}

type AppInfo struct {
	//mapstructure是什么：mapstructure是Viper支持的配置解析器，它可以将配置信息解析到结构体中
	//名字和yaml中的名字对应起来
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Port      int    `mapstructure:"port"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Path       string `mapstructure:"path"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"database"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("./config/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return //
	}
	// 把读取到的配置信息反序列化到 Conf 变量中
	// 什么是反序列化：把二进制数据转换为指定的数据结构
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
	// 监控配置文件变化
	viper.WatchConfig()
	// 当配置文件发生变更之后会调用的回调函数
	// 注意：通过viper.WatchConfig()方法监听配置文件的变更，当配置文件发生变更之后会调用的回调函数
	// OnConfigChange是一个回调函数，当配置文件发生变更之后会调用这个回调函数
	// 钩子函数：当配置文件发生变更之后会调用这个回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
		}
	})
	return
}
