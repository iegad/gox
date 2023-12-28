package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/iegad/gox/cerberus/m"
	dbcenter "github.com/iegad/gox/frm/db_center"
	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/web"
)

type userLoginReq struct {
	LoginType  int32  `json:"login_type"`
	ChannelID  int32  `json:"channel_id"`
	DeviceType byte   `json:"device_type"`
	UserID     int64  `json:"user_id"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
	VCode      string `json:"vcode"`
}

type userLoginRsp struct {
	UserInfo  *dbcenter.UserInfo `json:"user_info"`
	ProxyHost string             `json:"proxy_host"`
	ProxyCode string             `json:"proxy_code"`
	HallCode  string             `json:"hall_code"`
}

func UserLogin(c *gin.Context) {
	req := &userLoginReq{}
	err := c.BindJSON(req)
	if err != nil {
		log.Error(err)
		web.Response(c, -1, "request is invalid", nil)
		return
	}

	if req.LoginType <= 0 {
		web.Response(c, -1, "login_type is invalid", nil)
		return
	}

	if req.ChannelID <= 0 {
		web.Response(c, -1, "channel_id is invalid", nil)
		return
	}

	if req.DeviceType <= 0 {
		web.Response(c, -1, "device_type is invalid", nil)
		return
	}

	var uinfo *dbcenter.UserInfo
	switch req.LoginType {
	case 1: // 游客登录
		uinfo, err = guestLogin(req)

	case 2: // 用户名密码登录
		uinfo, err = userPasswdLogin(req)

	case 3: // 手机号验证码登录
		uinfo, err = usernameVCodeLogin(req)

	case 4: // 自动登录
		uinfo, err = autoLogin(req)
	}

	if err != nil {
		web.Response(c, -1, err.Error(), nil)
		return
	}

	err = uinfo.SetSession(m.Redis)
	if err != nil {
		log.Error(err)
		web.Response(c, -1, "set session failed", nil)
		return
	}

	// TODO: 记录登录日志
	hall := m.NodeManager.GetHall(req.ChannelID)
	proxy := m.NodeManager.GetProxy(req.ChannelID)

	var host string
	if req.DeviceType == 1 {
		host = proxy.FrontWsHost
	} else {
		host = proxy.FrontTcpHost
	}

	web.Response(c, 0, "", &userLoginRsp{
		UserInfo:  uinfo,
		ProxyHost: host,
		ProxyCode: proxy.NodeCode,
		HallCode:  hall.NodeCode,
	})
}

func guestLogin(req *userLoginReq) (*dbcenter.UserInfo, error) {
	uinfo, err := dbcenter.NewGuest(req.ChannelID, m.Mysql)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = uinfo.Insert(m.Mysql)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return uinfo, nil
}

func userPasswdLogin(req *userLoginReq) (*dbcenter.UserInfo, error) {
	return nil, nil
}

func usernameVCodeLogin(req *userLoginReq) (*dbcenter.UserInfo, error) {
	return nil, nil
}

func autoLogin(req *userLoginReq) (*dbcenter.UserInfo, error) {
	return nil, nil
}
