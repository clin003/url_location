package config

import (
	"fmt"
	"path/filepath"
	"url_location/config/config_embed"

	"gitee.com/lyhuilin/log"
	"gitee.com/lyhuilin/util"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	//初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}
	//初始化日志包
	c.initLog()
	//监控配置文件变化并热加载程序
	c.watchConfig()
	return nil
}

func (c *Config) initConfig() error {
	fileName := ""
	if c.Name != "" {
		viper.SetConfigFile(c.Name) //如果指定了配置文件，则解析指定的配置文件
		fileName = c.Name
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
		fileName = filepath.Join("conf", "config.yaml")
	}
	viper.SetConfigType("yaml") //设置配置文件格式为YAML

	// 检测是否存在配置文件，如果不存在，则自动生成
	if !util.FileExist(fileName) {
		fmt.Printf("配置文件(%s)不存在，自动生成默认配置\n", fileName)
		fileDir := filepath.Dir(fileName)
		if !util.IsExist(fileDir) {
			if err := util.MkDir(fileDir); err != nil {
				return fmt.Errorf("生成配置文件目录出错:%w", err)
			}
		}
		if err := util.WriteFile(fileName, config_embed.ConfigSimpleYaml); err != nil {
			return fmt.Errorf("生成配置文件出错:%w", err)
		}
	}

	// viper.AutomaticEnv() //读取匹配的环境变量
	// viper.SetEnvPrefix("HLTYAPI") //读取环境变量的前缀
	// replacer := strings.NewReplacer(".", "_")
	// viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { //viper 解析配置文件
		return err
	}

	return nil
}

// 日志配置加载
func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("llog.og_backup_count"),
	}
	log.InitWithConfig(&passLagerCfg)
}

//监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
}
