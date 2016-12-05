// Package cl provides a simple goroutine-safe capped list implementation.
package cl

import (
	"container/list"
	"sync"
)

// CappedList is a capped list where old items are evicted to make room for new
// items.
type CappedList struct {
	limit  int
	list   *list.List
	lookup map[interface{}]*list.Element
	mu     sync.Mutex
}

// Key is any type that is comparable.
type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

// New creates a new CappedList with the given item limit.
func New(limit int) *CappedList {
	return &CappedList{
		limit:  limit,
		list:   list.New(),
		lookup: make(map[interface{}]*list.Element),
	}
}

// All return all elements in the CappedList.
func (cl *CappedList) All() map[Key]interface{} {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	res := make(map[Key]interface{})
	for k, e := range cl.lookup {
		res[k] = e
	}

	return res
}

// Add adds an item to the CappedList with the given key and value.
func (cl *CappedList) Add(key Key, value interface{}) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	if e, ok := cl.lookup[key]; ok {
		cl.list.MoveToFront(e)
		e.Value.(*entry).value = value
		return
	}

	e := cl.list.PushFront(&entry{key, value})
	cl.lookup[key] = e
	cl.trim()
}

// Get returns an item from the list, returning the value and a bool
// representing whether or not it was found.
func (cl *CappedList) Get(key Key) (interface{}, bool) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	if e, ok := cl.lookup[key]; ok {
		e2 := e.Value.(*entry)
		return e2.value, ok
	}

	return nil, false
}

// Test determines whether the CappedList contains an item, returning true or false.
func (cl *CappedList) Contains(key Key) bool {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	if _, ok := cl.lookup[key]; ok {
		return true
	}

	return false
}

// Len returns the number of entries in the CappedList.
func (cl *CappedList) Len() int {
	return cl.list.Len()
}

// Limit returns the capacity of the CappedList.
func (cl *CappedList) Limit() int {
	return cl.limit
}

// Remove removes an item from the CappedList by key.
func (cl *CappedList) Remove(key Key) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	if e, ok := cl.lookup[key]; ok {
		cl.remove(e)
	}
}

func (cl *CappedList) trim() {
	for cl.limit != 0 && cl.list.Len() > cl.limit {
		if e := cl.list.Back(); e != nil {
			cl.remove(e)
		}
	}
}

func (cl *CappedList) remove(e *list.Element) {
	cl.list.Remove(e)
	kv := e.Value.(*entry)
	delete(cl.lookup, kv.key)
}
