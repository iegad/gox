package front

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/f"
)

func Info(c *gin.Context) {
	c.JSON(200, web.NewResponse(0, "", f.Service.Info()))
}
