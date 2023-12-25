package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/b"
)

func Info(c *gin.Context) {
	web.Response(c, 0, "", b.Service.Info())
}
