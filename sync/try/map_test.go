package try

import (
	"testing"
)

func TestMapGet(t *testing.T) {
	m := Map{}
	t1 := m.Get("key1")
	assert_equals(t, true, t1.Lock())
	assert_equals(t, false, t1.Lock())
	t1.Unlock()
	assert_equals(t, true, t1.Lock())
	t2 := m.Get("key2")
	assert_equals(t, true, t2.Lock())
	assert_equals(t, false, t2.Lock())
	t2.Unlock()
	assert_equals(t, true, t2.Lock())
}
