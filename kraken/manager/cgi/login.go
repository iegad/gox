package cgi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/web"
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

	if len(req.UserName) == 0 || len(req.UserName) < 5 || len(req.UserName) > 16 {
		c.JSON(http.StatusOK, web.NewResponse(-1, "user_name is invalid", nil))
		return
	}

	if len(req.Password) != 32 {
		c.JSON(http.StatusOK, web.NewResponse(-1, "password is invalid", nil))
		return
	}

	// TODO: authentication
	c.JSON(http.StatusOK, web.NewResponse(0, "", &LoginRsp{UserID: 1, Token: "11111111111111111111111111111111"}))
}