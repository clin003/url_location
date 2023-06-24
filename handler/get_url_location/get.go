package get_url_location

import (
	"strings"
	"url_location/get_env"
	"url_location/pkg/chromedp_location"
	"url_location/sdk"

	"gitee.com/lyhuilin/pkg/errno"
	"gitee.com/lyhuilin/pkg/handler"
	"gitee.com/lyhuilin/util"
	"github.com/gin-gonic/gin"
)

type UrlLocationRequest struct {
	UrlText string `json:"url"  form:"url"` //origin_url
	// ShareSource string `json:"shareSource"  form:"shareSource"`
	// BotGID      string `json:"botGID"  form:"botGID"` //机器人消息groupid
	// BotID       string `json:"botID"  form:"botID"`   //机器人id
	sdk.BaseRequest
	// IsJson  bool   `json:"retJson"  form:"retJson"` //是否返回Json
	// Botoken string `json:"botoken"  form:"botoken"` //机器人Token
}

func HandlerGetUrlOrigin(c *gin.Context) {
	var r UrlLocationRequest
	if err := c.ShouldBind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if r.UrlText == "" {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	urlText := r.UrlText
	urlText = strings.Join(strings.Fields(urlText), "")
	retUrl, err := getUrlOriginUrl(urlText)
	handler.SendResponse(c, err, map[string]string{
		"origin_url": retUrl,
		// "location_url": retUrl,
		"url": urlText,
	})
}

func HandlerGetUrlLocation(c *gin.Context) {
	var r UrlLocationRequest
	if err := c.ShouldBind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if r.UrlText == "" {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	urlText := r.UrlText
	urlText = strings.Join(strings.Fields(urlText), "")
	retUrl, err := getUrlLocation(urlText)
	handler.SendResponse(c, err, map[string]string{
		// "origin_url":   message,
		"location_url": retUrl,
		"url":          urlText,
	})
}
func HandlerGetUrlOriginAndLocation(c *gin.Context) {
	var r UrlLocationRequest
	if err := c.ShouldBind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if r.UrlText == "" {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	urlText := r.UrlText
	urlText = strings.Join(strings.Fields(urlText), "")
	originUrl, locationUrl, err := getUrlLocationAndOriginUrl(urlText)
	handler.SendResponse(c, err, map[string]string{
		"origin_url":   originUrl,
		"location_url": locationUrl,
		"url":          urlText,
	})
}

func getUrlLocationAndOriginUrl(urlStr string) (originUrl, locationUrl string, err error) {
	locationUrl, err = getUrlLocation(urlStr)
	originUrl, err = getUrlOriginUrl(urlStr)
	return originUrl, locationUrl, err
}
func getUrlLocation(urlStr string) (string, error) {
	locationUrl, err := util.GetRedirectUrl(urlStr, false)
	if err != nil {
		locationUrl, err = util.GetRedirectUrlEx(urlStr, false)
	}

	if strings.EqualFold(locationUrl, urlStr) || err != nil {
		locationUrl = ""
	}
	return locationUrl, err
}
func getUrlOriginUrl(urlStr string) (string, error) {
	remoteDebuggingUrl := get_env.GetChromedpRemoteDebuggingUrl()
	remoteDebuggingPort := get_env.GetChromedpRemoteDebuggingPort()
	// fmt.Println("getUrlOriginUrl", remoteDebuggingUrl, remoteDebuggingPort)
	return chromedp_location.GetUrlLocation(urlStr, remoteDebuggingUrl, remoteDebuggingPort, 1)
}
