package m

import (
	"encoding/binary"
	"encoding/hex"
	"math/rand"

	"github.com/iegad/gox/frm/biz"
	"github.com/iegad/gox/frm/utils"
)

var (
	Players biz.UserManager
	Nodes   NodeManager
)

const PWD = "1234567890"

var Admin *admin

type admin struct {
	UserName   string
	Password   string
	Token      string
	Idempotent int64
	IPAddr     string
}

func Init() {
	Admin = &admin{
		UserName: "root",
		Password: utils.MD5Hex(PWD),
	}
}

func (this_ *admin) Login(ipAddr string) string {
	rint := rand.Uint64()
	tmp := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp, rint)
	this_.Token = hex.EncodeToString(tmp)
	this_.IPAddr = ipAddr
	this_.Idempotent = 0
	return this_.Token
}
