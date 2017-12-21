package oldpkg

import (
	"testing"

	"github.com/hyper0x/review2017-go/newpkg"
)

func TestOldType(t *testing.T) {
	var old OldType
	switch tt := interface{}(old).(type) {
	case newpkg.NewType:
		t.Logf("The type of 'old' is equals to '%T'.", old)
	default:
		t.Errorf("The type of 'old' does not conform to the rules: %T", tt)

	}
}
