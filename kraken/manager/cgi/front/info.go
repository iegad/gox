package front

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/f"
	"github.com/iegad/gox/kraken/m"
)

type InfoReq struct {
	web.BasicRequest
}

func Info(c *gin.Context) {
	req := &InfoReq{}
	err := c.BindJSON(req)
	if err != nil {
		log.Error(err)
		return
	}

	if req.UserID == nil || *req.UserID != 1 {
		c.JSON(200, web.NewResponse(-1, "user_id is invalid", nil))
		return
	}

	if req.Token == nil || len(*req.Token) != 16 || *req.Token != m.Admin.Token {
		c.JSON(200, web.NewResponse(-1, "token is invalid", nil))
		return
	}

	if req.Idempotent == nil || *req.Idempotent <= m.Admin.Idempotent {
		c.JSON(200, web.NewResponse(-1, "idemponent is invalid", nil))
		return
	}

	m.Admin.Idempotent = *req.Idempotent
	c.JSON(200, web.NewResponse(0, "", f.Service.Info()))
}
