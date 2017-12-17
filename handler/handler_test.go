package handler

import (
	"testing"
	"github.com/hyper0x/review2017-go/oldpkg"
	"github.com/hyper0x/review2017-go/newpkg"
)

func TestHandleXXX(t *testing.T) {
	HandleXXX(newpkg.NewType(1))
	HandleXXX(oldpkg.OldType(0))
}