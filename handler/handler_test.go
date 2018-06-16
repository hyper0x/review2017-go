package handler

import (
	"github.com/hyper0x/review2017-go/newpkg"
	"github.com/hyper0x/review2017-go/oldpkg"
	"testing"
)

func TestHandleXXX(t *testing.T) {
	HandleXXX(newpkg.NewType(1))
	HandleXXX(oldpkg.OldType(0))
}
