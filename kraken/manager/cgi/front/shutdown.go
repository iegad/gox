package front

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/f"
)

func Shutdown(c *gin.Context) {
	f.Service.Shutdown()
	c.JSON(200, web.NewResponse(0, "", nil))
}
