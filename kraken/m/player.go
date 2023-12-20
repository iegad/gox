package m

import (
	"encoding/binary"
	"math/rand"

	"github.com/iegad/gox/frm/nw"
)

type player struct {
	UserID int64
	Token  []byte
	sess   *nw.Sess
}

func NewPlayer(userID int64, sess *nw.Sess) *player {
	rint := rand.Uint64()

	tmp := &player{
		UserID: userID,
		sess:   sess,
	}

	sess.UserData = tmp
	binary.BigEndian.PutUint64(tmp.Token, rint)
	return tmp
}

func (this_ *player) Session() *nw.Sess {
	return this_.sess
}
