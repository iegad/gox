package front

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/ios"
)

func Shutdown(c *gin.Context) {
	ios.Front.Shutdown()
	c.JSON(200, web.NewResponse(0, "", nil))
}