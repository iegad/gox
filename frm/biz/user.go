package biz

import (
	"encoding/binary"
	"log"
	"math/rand"
	"sync"

	"github.com/iegad/gox/frm/nw"
)

type user struct {
	UserID int64
	Token  []byte
	sess   *nw.Sess
}

func NewUser(userID int64, sess *nw.Sess) *user {
	rint := rand.Uint64()

	tmp := &user{
		UserID: userID,
		sess:   sess,
	}

	sess.UserData = tmp
	binary.BigEndian.PutUint64(tmp.Token, rint)
	return tmp
}

func (this_ *user) Session() *nw.Sess {
	return this_.sess
}

type UserManager struct {
	userMap sync.Map
}

func (this_ *UserManager) AddPlayer(plr *user) error {
	if plr == nil || plr.sess == nil || plr.Token == nil || plr.UserID <= 0 {
		log.Fatal("plr is invalid")
	}

	if _, ok := this_.userMap.Load(plr.UserID); ok {
		return Err_PlayerIsExists
	}

	this_.userMap.Delete(plr.sess.RemoteAddr().String())
	this_.userMap.Store(plr.UserID, plr)
	return nil
}

func (this_ *UserManager) AddSession(sess *nw.Sess) {
	if sess == nil || sess.UserData != nil {
		log.Fatal("sess is nil")
	}

	this_.userMap.Store(sess.RemoteAddr().String(), sess)
}

func (this_ *UserManager) GetUserByID(userID int64) *user {
	if userID <= 0 {
		log.Fatal("userID is invalid")
	}

	if v, ok := this_.userMap.Load(userID); ok {
		return v.(*user)
	}

	return nil
}

func (this_ *UserManager) GetSession(remoteAddr string) *nw.Sess {
	if len(remoteAddr) == 0 {
		log.Fatal("remoteAddr is invalid")
	}

	if v, ok := this_.userMap.Load(remoteAddr); ok {
		return v.(*nw.Sess)
	}

	return nil
}

func (this_ *UserManager) RemoveSession(remoteAddr string) {
	if len(remoteAddr) == 0 {
		log.Fatal("remoteAddr is invalid")
	}

	if sess, ok := this_.userMap.LoadAndDelete(remoteAddr); ok {
		sess.(*nw.Sess).Shutdown()
	}
}

func (this_ *UserManager) RemovePlayer(userID int64) {
	if userID <= 0 {
		log.Fatal("userID is invalid")
	}

	if plr, ok := this_.userMap.LoadAndDelete(userID); ok {
		plr.(*user).sess.Shutdown()
	}
}

func (this_ *UserManager) Clear() {
	this_.userMap.Range(func(key, value any) bool {
		if _, ok := key.(string); ok {
			value.(*nw.Sess).Shutdown()
		} else {
			value.(*user).sess.Shutdown()
		}

		return true
	})

	this_.userMap = sync.Map{}
}
