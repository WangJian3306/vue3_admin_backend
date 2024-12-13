package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

// Conf 全局变量，用来保存程序的所有配置信息
var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*SnowFlake   `mapstructure:"snowflake"`
	*Static      `mapstructure:"static"`
}

type Static struct {
	Host string `mapstructure:"host"`
	Path string `mapstructure:"path"`
}

type SnowFlake struct {
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func Init(configPath string) (err error) {
	//viper.SetConfigName("config") // 指定配置文件名称（不需要带后缀）
	//viper.SetConfigType("yaml") // 指定配置文件类型（专用于从远程获取配置信息时指定配置文件类型的）
	viper.SetConfigFile(configPath)
	viper.AddConfigPath(".")                  // 指定查找文件的路径（配合相对路径使用）
	viper.AutomaticEnv()                      // 指定支持从环境变量读取配置
	viper.SetEnvPrefix("VUE3_ADMIN")          // 指定环境变量 KEY 的前缀
	replacer := strings.NewReplacer(".", "_") // 替换规则，将 . 替换为 _
	viper.SetEnvKeyReplacer(replacer)
	err = viper.ReadInConfig()
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err: %v\n", err)
		return err
	}
	// 把读取到的配置信息，反序列化到 Conf 变量中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v \n", err)
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v \n", err)
		}
	})

	return err
}
