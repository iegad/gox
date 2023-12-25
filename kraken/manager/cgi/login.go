package cgi

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/web"
	"github.com/iegad/gox/kraken/m"
)

type LoginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginRsp struct {
	UserID int32  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

func Login(c *gin.Context) {
	req := &LoginReq{}
	err := c.BindJSON(req)
	if err != nil {
		log.Error(err)
		return
	}

	if len(req.UserName) == 0 || req.UserName != m.Admin.UserName {
		web.Response(c, -1, "user_name is invalid", nil)
		return
	}

	if len(req.Password) != 32 || m.Admin.Password != req.Password {
		web.Response(c, -1, "password is invalid", nil)
		return
	}

	token := m.Admin.Login(c.RemoteIP())
	web.Response(c, 0, "", &LoginRsp{UserID: 1, Token: token})
}
