package front

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/f"
	"github.com/iegad/gox/kraken/m"
)

func Run(c *gin.Context) {
	req := &web.BasicRequest{}
	err := c.BindJSON(req)
	if err != nil {
		log.Error(err)
		return
	}

	if req.UserID == nil || *req.UserID != 1 {
		web.Response(c, -1, "user_id is invalid", nil)
		return
	}

	if req.Token == nil || len(*req.Token) != 16 || *req.Token != m.Admin.Token {
		web.Response(c, -1, "token is invalid", nil)
		return
	}

	if req.Idempotent == nil || *req.Idempotent <= m.Admin.Idempotent {
		web.Response(c, -1, "idempotent is invalid", nil)
		return
	}

	m.Admin.Idempotent = *req.Idempotent

	err = f.Service.Run()
	if err != nil {
		log.Error(err)
		web.Response(c, -1, err.Error(), nil)
		return
	}

	web.Response(c, 0, "", nil)
}
