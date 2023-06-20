package main

import (
	"fmt"
	"time"
	"url_location/pkg/constvar"

	"gitee.com/lyhuilin/log"
	"gitee.com/lyhuilin/util"
	"github.com/spf13/viper"
)

// 自检openAPI服务是否正常运行
func pingServer() error {
	apiURL := viper.GetString("url")
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {

		if util.CheckPingBaseURL(apiURL) {
			return nil
		}

		log.Infof("(%s)等待自检, 1秒后重试(%d) %s", constvar.APP_NAME, i, apiURL)
		time.Sleep(time.Second)
	}
	return fmt.Errorf("(%s)自检失败 %s", constvar.APP_NAME, apiURL)
}
