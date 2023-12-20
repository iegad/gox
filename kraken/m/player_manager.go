package m

import (
	"sync"

	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/nw"
	"github.com/iegad/gox/kraken/basic"
)

type PlayerManager struct {
	playerMap sync.Map
}

func (this_ *PlayerManager) AddPlayer(plr *player) error {
	if plr == nil || plr.sess == nil || plr.Token == nil || plr.UserID <= 0 {
		log.Fatal("plr is invalid")
	}

	if _, ok := this_.playerMap.Load(plr.UserID); ok {
		return basic.Err_F_PlayerIsExists
	}

	this_.playerMap.Delete(plr.sess.RemoteAddr().String())
	this_.playerMap.Store(plr.UserID, plr)
	return nil
}

func (this_ *PlayerManager) AddSession(sess *nw.Sess) {
	if sess == nil || sess.UserData != nil {
		log.Fatal("sess is nil")
	}

	this_.playerMap.Store(sess.RemoteAddr().String(), sess)
}

func (this_ *PlayerManager) GetUserByID(userID int64) *player {
	if userID <= 0 {
		log.Fatal("userID is invalid")
	}

	if v, ok := this_.playerMap.Load(userID); ok {
		return v.(*player)
	}

	return nil
}

func (this_ *PlayerManager) GetSession(remoteAddr string) *nw.Sess {
	if len(remoteAddr) == 0 {
		log.Fatal("remoteAddr is invalid")
	}

	if v, ok := this_.playerMap.Load(remoteAddr); ok {
		return v.(*nw.Sess)
	}

	return nil
}

func (this_ *PlayerManager) RemoveSession(remoteAddr string) {
	if len(remoteAddr) == 0 {
		log.Fatal("remoteAddr is invalid")
	}

	if sess, ok := this_.playerMap.LoadAndDelete(remoteAddr); ok {
		sess.(*nw.Sess).Shutdown()
	}
}

func (this_ *PlayerManager) RemovePlayer(userID int64) {
	if userID <= 0 {
		log.Fatal("userID is invalid")
	}

	if plr, ok := this_.playerMap.LoadAndDelete(userID); ok {
		plr.(*player).sess.Shutdown()
	}
}

func (this_ *PlayerManager) Clear() {
	this_.playerMap.Range(func(key, value any) bool {
		if _, ok := key.(string); ok {
			value.(*nw.Sess).Shutdown()
		} else {
			value.(*player).sess.Shutdown()
		}

		return true
	})

	this_.playerMap = sync.Map{}
}
