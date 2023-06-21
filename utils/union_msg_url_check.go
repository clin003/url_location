package utils

import (
	"strings"
	"url_location/get_env"

	"gitee.com/lyhuilin/util"
)

//检查Url是否支持可转链，true 支持，false 不支持
func UnionMsgUrlCheck(msgContext string) bool {
	urlText := strings.ToLower(msgContext)
	if !strings.HasPrefix(urlText, "http") {
		return false
	}
	domainStr := get_env.GetUnionMsgDomain()
	domainList := util.KeyworldListParse(domainStr)
	for _, v := range domainList {
		vc := v
		if strings.EqualFold(urlText, vc) || strings.Contains(urlText, vc) {
			return true
		}
	}
	return false
}

func CanWaitUrlCheck(msgContext string) bool {
	urlText := strings.ToLower(msgContext)
	if !strings.HasPrefix(urlText, "http") {
		return false
	}
	domainStr := get_env.GetCanWaitDomain()
	domainList := util.KeyworldListParse(domainStr)
	for _, v := range domainList {
		vc := v
		if strings.EqualFold(urlText, vc) || strings.Contains(urlText, vc) {
			return true
		}
	}
	return false
}
