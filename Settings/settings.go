package Settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config = new(Appconfig)

type Appconfig struct {
	Name             string `mapstructure:"name"`
	Mode             string `mapstructure:"mode"`
	Version          string `mapstructure:"version"`
	Port             int    `mapstructure:"port"`
	*LogConfig       `mapstructure:"log"`
	*MysqlConfig     `mapstructure:"mysql"`
	*RedisConfig     `mapstructure:"redis"`
	*RabbitMQConfig  `mapstructure:"rabbit_mq"`
	*SnowFlakeConfig `mapstructure:"snowflake"`
}
type SnowFlakeConfig struct {
	MachineID int64  `mapstructure:"machine_id"`
	StartTime string `mapstructure:"start_time"`
}
type RabbitMQConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"user_name"`
	Password string `mapstructure:"password"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	UserName     string `mapstructure:"user_name"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (error error) {
	viper.SetConfigName("config")      // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")        // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./settings/") // 查找配置文件所在的路径
	err := viper.ReadInConfig()        // 查找并读取配置文件
	if err != nil {                    // 处理读取配置文件的错误
		fmt.Printf("Fatal error config file: %s \n", err)
		return err
	}
	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(&Config); err != nil {
			fmt.Printf("Fatal error config file: %s \n", err)
		}
	})

	fmt.Println("Config file loaded successfully")
	return nil
}
