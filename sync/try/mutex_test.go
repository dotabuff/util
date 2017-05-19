package try

import (
	"testing"
)

func TestMutex(t *testing.T) {
	m := Mutex{}
	assert_equals(t, true, m.Lock())
	assert_equals(t, false, m.Lock())
	m.Unlock()
	assert_equals(t, true, m.Lock())
	m.Unlock()

	panicked := false
	func() {
		defer func() {
			if err := recover(); err != nil {
				panicked = true
			}
		}()
		m.Unlock()
	}()
	assert_equals(t, true, panicked)
}
