package main

import (
	"fmt"
	"net/http"
	"path"
	"time"
	"url_location/config"
	"url_location/pkg/constvar"
	"url_location/router"

	"gitee.com/lyhuilin/log"
	"gitee.com/lyhuilin/pkg/env"
	"gitee.com/lyhuilin/util"
	"gitee.com/lyhuilin/version"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/crypto/acme/autocert"
)

var (
	cfg        = pflag.StringP("config", "c", "", "hltyapi config file path")
	versionArg = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()
	if *versionArg {
		fmt.Println(version.APPVersionEx())
		return
	}
	fmt.Printf("%s(%v) \n%s\n", constvar.APP_NAME, version.APPVersionEx(), constvar.APPDesc())
	fmt.Printf("(%s)运行在：%s\n", constvar.APP_NAME, env.Getwd())
	time.Sleep(time.Second * 3)

	// 加载配置文件
	if err := config.Init(*cfg); err != nil {
		fmt.Printf("(%s)加载配置文件失败:%v", constvar.APP_NAME, err)
		panic(err)
	}

	Init()
	ginMode := "release"
	if viper.GetString("runmode") != "" {
		ginMode = viper.GetString("runmode")
	}
	gin.SetMode(ginMode)
	g := gin.New()
	router.Load(
		g,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatalf(err, "(%s)没有响应，请检查配置及网络状态.", constvar.APP_NAME)
		}
		log.Infof("(%s)成功部署，服务地址:%s", constvar.APP_NAME, viper.GetString("url"))
	}()

	// 绑定开启服务端口
	if viper.GetBool("autotls_enable") {
		pemFile := viper.GetString("tls_pem_file")
		keyFile := viper.GetString("tls_key_file")
		hasPem, _ := util.IsExists(pemFile)
		hasKey, _ := util.IsExists(keyFile)
		if hasPem && hasKey {
			go func() {
				log.Info(http.ListenAndServeTLS(":443", path.Clean(pemFile), path.Clean(keyFile), g).Error())
			}()
		} else {
			go func() {
				m := autocert.Manager{
					Prompt:     autocert.AcceptTOS,
					HostPolicy: autocert.HostWhitelist(viper.GetString("autotls_domain")),
					Cache:      autocert.DirCache("autocert"),
				}
				log.Info(autotls.RunWithManager(g, &m).Error())
			}()
		}
	}

	if err := http.ListenAndServe(viper.GetString("addr"), g); err != nil {
		log.Errorf(err, "(%s)出错了，需要重启！", constvar.APP_NAME)
	}
}
