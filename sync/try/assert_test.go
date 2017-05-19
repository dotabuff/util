package try

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, condition bool, msg string, v ...interface{}) {
	if !condition {
		t.Errorf(msg, v...)
	}
}

func assert_equals(t *testing.T, exp, act interface{}) {
	assert(t, reflect.DeepEqual(exp, act), "expected %+v, got %+v", exp, act)
}
