package cl

import (
	"reflect"
	"testing"
)

func TestAddition(t *testing.T) {
	cl := New(3)
	cl.Add("jason", 3.14159)

	// Retains added items
	assert_equals(t, cl.Contains("jason"), true)
	assert_equals(t, cl.Len(), 1)

	// Finds added items
	val, found := cl.Get("jason")
	assert_equals(t, val, 3.14159)
	assert_equals(t, found, true)

	// Adds additional items and keeps count
	cl.Add("david", 1.23456)
	assert_equals(t, cl.Contains("david"), true)
	assert_equals(t, cl.Len(), 2)
}

func TestRemoval(t *testing.T) {
	cl := New(3)
	cl.Add("james", 1)
	cl.Add("jason", 1)
	cl.Add("joseph", 1)

	cl.Remove("jason")

	// Removes intended item
	assert_equals(t, cl.Contains("jason"), false)

	// Does not remove other items
	assert_equals(t, cl.Contains("james"), true)
	assert_equals(t, cl.Contains("joseph"), true)

	// Keeps a proper count
	assert_equals(t, cl.Len(), 2)
}

func TestEviction(t *testing.T) {
	cl := New(2)
	cl.Add("james", 1)
	cl.Add("jason", 2)
	cl.Add("joseph", 3)
	cl.Add("justin", 4)

	// Trims all internal lists
	assert_equals(t, cl.Len(), 2)
	assert_equals(t, cl.list.Len(), 2)
	assert_equals(t, len(cl.lookup), 2)

	// Retains newest items
	assert_equals(t, cl.Contains("justin"), true)
	assert_equals(t, cl.Contains("joseph"), true)

	// Discards oldest items
	assert_equals(t, cl.Contains("jason"), false)
	assert_equals(t, cl.Contains("james"), false)
}

func assert(t *testing.T, condition bool, msg string, v ...interface{}) {
	if !condition {
		t.Errorf(msg, v...)
	}
}

func assert_equals(t *testing.T, exp, act interface{}) {
	assert(t, reflect.DeepEqual(exp, act), "expected %+v, got %+v", exp, act)
}
