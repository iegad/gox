package front

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/f"
)

func Run(c *gin.Context) {
	err := f.Service.Run()
	if err != nil {
		log.Error(err)
		web.Response(c, -1, err.Error(), nil)
		return
	}

	web.Response(c, 0, "", nil)
}
