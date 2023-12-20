package front

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/f"
)

func Run(c *gin.Context) {
	code := int32(0)

	err := f.Service.Run()
	if err != nil {
		log.Error(err)
		code = -1
	}

	c.JSON(200, web.NewResponse(code, err.Error(), nil))
}
