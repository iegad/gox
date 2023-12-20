package m

import (
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/iegad/gox/frm/log"
	"github.com/iegad/gox/frm/utils"
)

var Admin *admin

type admin struct {
	UserName   string
	Password   string
	Token      string
	Idempotent int64
}

func Init() {
	Admin = &admin{
		UserName: "admin",
	}

	go func() {
		for {
			pwd := utils.RandPassword(16, true)
			log.Info("%v: %v", pwd, len(pwd))
			Admin.Password = pwd
			interval := rand.Int63n(6) + 1
			time.Sleep(time.Hour * time.Duration(interval))
		}
	}()
}

func (this_ *admin) GenerateToken() string {
	rint := rand.Uint64()
	tmp := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp, rint)
	this_.Token = hex.EncodeToString(tmp)
	return this_.Token
}
