package get_env

import (
	// "fmt"
	"os"

	"github.com/spf13/viper"
)

func GetChromedpRemoteDebuggingUrl() string {
	tmpText := viper.GetString("chromedp_remote_debug_url")
	if len(tmpText) > 0 {
		// fmt.Println("GetChromedpRemoteDebuggingUrl", tmpText)
		return tmpText
	}
	tmpText = os.Getenv("CHROMEDP_REMOTE_DEBUG_URL")
	if len(tmpText) > 0 {
		// fmt.Println("GetChromedpRemoteDebuggingUrl", tmpText)
		return tmpText
	}
	return ""
	// return os.Getenv("CHROMEDP_REMOTE_DEBUG_URL")
}
func GetChromedpRemoteDebuggingPort() string {
	tmpText := viper.GetString("chromedp_remote_debug_port")
	if len(tmpText) > 0 {
		// fmt.Println("GetChromedpRemoteDebuggingUrl", tmpText)
		return tmpText
	}
	tmpText = os.Getenv("CHROMEDP_REMOTE_DEBUG_PORT")
	if len(tmpText) > 0 {
		// fmt.Println("GetChromedpRemoteDebuggingUrl", tmpText)
		return tmpText
	}
	return "9222"
	// return os.Getenv("CHROMEDP_REMOTE_DEBUG_URL")
}

// func GetCanWaitDomain() string {
// 	retText := "https://www.zzzdm.com/go|https://go.smzdm.com"
// 	return retText
// }
func GetCanWaitDomain() string {
	tmpText := viper.GetString("chromedp_can_wait_domain_list")
	if len(tmpText) > 0 {
		return tmpText
	}
	tmpText = os.Getenv("CHROMEDP_CAN_WAIT_DOMAIN_LIST")
	if len(tmpText) > 0 {
		return tmpText
	}
	return "https://www.zzzdm.com/go|https://go.smzdm.com"
}
