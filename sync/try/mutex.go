package try

import (
	"sync/atomic"
)

// Mutex provides a mutual exclusion lock. It differs from the Mutex provided
// by the sync package in that Lock() does not block and instead returns a
// boolean indicating whether or not the lock was successfully obtained.
type Mutex struct {
	r uint32
}

// Lock locks the mutex, returning true if successful.
func (m *Mutex) Lock() bool {
	return atomic.CompareAndSwapUint32(&m.r, 0, 1)
}

// Unlock unlocks a previously locked Mutex.
//
// Unlock will panic if the Mutex was not previously locked.
func (m *Mutex) Unlock() {
	if !atomic.CompareAndSwapUint32(&m.r, 1, 0) {
		panic("try: unlock of unlocked Mutex")
	}
}
