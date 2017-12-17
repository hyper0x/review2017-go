package oldpkg

import (
	"testing"
	"github.com/hyper0x/review2017-go/newpkg"
)

func TestOldType(t *testing.T) {
	var old OldType
	switch tt := interface{}(old).(type) {
	case newpkg.NewType:
		t.Logf("The type of 'old' is equal to '%T'.", old)
	default:
		t.Errorf("The type of old '%s' is not conform to the rules", tt)

	}
}


