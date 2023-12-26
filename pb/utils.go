package pb

import (
	"github.com/iegad/gox/frm/log"
	"google.golang.org/protobuf/proto"
)

func CheckNodePackage(pack *Package) error {
	if pack == nil {
		log.Fatal("pack is nil")
	}

	if len(pack.NodeCode) != 16 {
		return Err_NodeCodeInvalid
	}

	if pack.MessageID <= 0 {
		return Err_MessageIDInvalid
	}

	if pack.Idempotent <= 0 {
		return Err_IdempotentInvalid
	}

	return nil
}

func SerializeNodePackage(nodeCode []byte, mid int32, idempotent int64, data []byte) ([]byte, error) {
	return proto.Marshal(&Package{
		NodeCode:   nodeCode,
		MessageID:  mid,
		Idempotent: idempotent,
		Data:       data,
	})
}

func ParseNodePackage(data []byte) (*Package, error) {
	pack := &Package{}
	err := proto.Unmarshal(data, pack)
	if err != nil {
		return nil, err
	}

	err = CheckNodePackage(pack)
	if err != nil {
		return nil, err
	}

	return pack, nil
}

type Msg[T any] interface {
	*T
	proto.Message
}

func ParseMessage[T any, U Msg[T]](data []byte) (*T, error) {
	msg := U(new(T))
	err := proto.Unmarshal(data, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
