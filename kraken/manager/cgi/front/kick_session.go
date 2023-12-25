package front

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/f"
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

	if len(req.RemoteAddr) == 0 {
		web.Response(c, -1, "remote_addr is invalid", nil)
		return
	}

	f.Service.Engine.PlayerManager.RemoveSession(req.RemoteAddr)
	web.Response(c, 0, "", nil)
}
