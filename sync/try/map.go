package try

import (
	"sync"
)

// Map provides a map of Mutex objects keyed by string. The zero value
// of a Map is safe for use.
type Map struct {
	mu sync.Mutex
	ts map[string]*Mutex
}

// Get returns the Mutex associated with a given key, creating one if it
// does not already exist. Get cam b
func (m *Map) Get(key string) *Mutex {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.ts == nil {
		m.ts = make(map[string]*Mutex)
	}
	if _, ok := m.ts[key]; !ok {
		m.ts[key] = &Mutex{}
	}
	return m.ts[key]
}
