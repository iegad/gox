package pb

import "github.com/iegad/gox/frm/log"

func CheckPackage(pack *Package) error {
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
