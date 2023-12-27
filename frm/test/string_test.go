package test

import (
	"testing"

	"github.com/iegad/gox/frm/utils"
)

func TestString(t *testing.T) {
	s := "Hello world"
	b := utils.Str2Bytes(s)
	sc := utils.Bytes2Str(b)

	t.Logf("s[%p], b[%p]\n", &s, &b)
	t.Logf("s[%p], sc[%p]\n", &s, &sc)
	t.Logf("%v, %v\n", s, sc)
}
