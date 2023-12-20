package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/b"
)

func Shutdown(c *gin.Context) {
	b.Service.Shutdown()
	c.JSON(200, web.NewResponse(0, "", nil))
}
