package front

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/m"
)

type KickSessionReq struct {
	web.BasicRequest
	RemoteAddr string `json:"remote_addr,omitempty"`
}

func KickSession(c *gin.Context) {
	req := &KickSessionReq{}
	err := c.BindJSON(req)
	if err != nil {
		web.Response(c, -1, err.Error(), nil)
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

	if len(req.RemoteAddr) == 0 {
		web.Response(c, -1, "remote_addr is invalid", nil)
		return
	}

	m.Players.RemoveSession(req.RemoteAddr)
	web.Response(c, 0, "", nil)
}
