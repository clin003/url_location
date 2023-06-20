package router

import (
	. "url_location/handler"
	"url_location/pkg/constvar"

	"gitee.com/lyhuilin/pkg/errno"
	"github.com/gin-gonic/gin"
)

// 404 Not found
func ApiNotFound(c *gin.Context) {
	SendResponse(c, errno.Err404, constvar.APPDesc404())
}

// API Hello
func ApiHello(c *gin.Context) {
	SendResponse(c, errno.SayHello, constvar.APPDesc())
}

// API ping
func ApiPing(c *gin.Context) {
	SendResponse(c, errno.PONG, constvar.APPVersionEx())
}
