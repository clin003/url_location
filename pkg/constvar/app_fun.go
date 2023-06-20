package constvar

import (
	"fmt"

	"time"

	"gitee.com/lyhuilin/version"
)

func APPDesc() string {
	return fmt.Sprintf("交流QQ群: 1051824036, 网站：www.lyhuilin.com \nCopyright ©2018-%d LYHUILIN Team. All Rights Reserved", time.Now().Year())
}
func APPDesc404() string {
	return fmt.Sprintf("慧林淘友交流QQ群：153690156 ，网站：www.lyhuilin.com (Copyright ©2018-%d LYHUILIN Team All Rights Reserved)(Error API route.)", time.Now().Year())
}

func APPDescEx() string {
	return fmt.Sprintf("openAPI(%s) 交流QQ群: 1051824036, 网站：www.lyhuilin.com \nCopyright ©2018-%d LYHUILIN Team. All Rights Reserved", version.APPVersion(), time.Now().Year())
}

func APPDevInfo() string {
	return version.APPDevInfo()
}
func APPVersion() string {
	return version.APPVersion()
}
func APPVersionEx() string {
	return version.APPVersionEx()
}

func IsDev() bool {
	return version.IsDev()
}
