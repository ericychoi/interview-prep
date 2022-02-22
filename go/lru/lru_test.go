package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLRU(t *testing.T) {
	l := NewLRU(3)
	l.Add("1", "a")
	assert.Equal(t, "1", l.tail.key)
	assert.Equal(t, "1", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"1"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"1": "a"}) })
	l.Dump()

	assert.Equal(t, "a", l.Read("1")) // read one element
	assert.Equal(t, "1", l.tail.key)
	assert.Equal(t, "1", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"1"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"1": "a"}) })
	l.Dump()

	assert.Equal(t, "", l.Read("?")) // read one element, non-existent
	assert.Equal(t, "1", l.tail.key)
	assert.Equal(t, "1", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"1"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"1": "a"}) })
	l.Dump()

	l.Remove("?") // remove one element, non-existent: 1 => 1
	assert.Equal(t, "1", l.tail.key)
	assert.Equal(t, "1", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"1"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"1": "a"}) })
	l.Dump()

	l.Remove("1") // remove one element : 1 => nil
	assert.Nil(t, l.tail)
	assert.Nil(t, l.head)
	assert.Empty(t, l.store)
	l.Dump()

	l.Add("1", "a")
	l.Add("2", "b")
	l.Add("3", "c")
	l.Add("4", "d") // evicted, 2,3,4
	assert.Equal(t, "4", l.tail.key)
	assert.Equal(t, "2", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"2", "3", "4"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"2": "b", "4": "d", "3": "c"}) })
	l.Dump()

	assert.Equal(t, "c", l.Read("3")) // 2,3,4 => 2,4,3
	assert.Equal(t, "3", l.tail.key)
	assert.Equal(t, "2", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"2", "4", "3"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"2": "b", "4": "d", "3": "c"}) })
	l.Dump()

	l.Add("5", "e") // 2,4,3 => 4,3,5
	assert.Equal(t, "5", l.tail.key)
	assert.Equal(t, "4", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"4", "3", "5"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"4": "d", "5": "e", "3": "c"}) })
	l.Dump()

	assert.Equal(t, "d", l.Read("4")) // read head: 4,3,5 => 3,5,4
	assert.Equal(t, "4", l.tail.key)
	assert.Equal(t, "3", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"3", "5", "4"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"4": "d", "5": "e", "3": "c"}) })
	l.Dump()

	l.Add("6", "f") // 3,5,4 => 5,4,6
	assert.Equal(t, "6", l.tail.key)
	assert.Equal(t, "5", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"5", "4", "6"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"6": "f", "4": "d", "5": "e"}) })
	l.Dump()

	assert.Equal(t, "f", l.Read("6")) // read tail: 5,4,6 => 5,4,6
	assert.Equal(t, "6", l.tail.key)
	assert.Equal(t, "5", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"5", "4", "6"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"6": "f", "4": "d", "5": "e"}) })
	l.Dump()

	assert.Equal(t, "", l.Read("1")) // read non-existent value: 5,4,6 => 5,4,6
	assert.Equal(t, "6", l.tail.key)
	assert.Equal(t, "5", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"5", "4", "6"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"6": "f", "4": "d", "5": "e"}) })
	l.Dump()

	l.Remove("4") // remove: 5,4,6 => 5,6
	assert.Equal(t, "6", l.tail.key)
	assert.Equal(t, "5", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"5", "6"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"6": "f", "5": "e"}) })
	l.Dump()

	l.Remove("6") // remove from tail: 5,6 => 5
	assert.Equal(t, "5", l.tail.key)
	assert.Equal(t, "5", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"5"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"5": "e"}) })
	l.Dump()

	// check eviction one more time
	l.Add("7", "g")
	l.Add("8", "h")
	l.Add("9", "i") // 5 => 7,8,9
	assert.Equal(t, "9", l.tail.key)
	assert.Equal(t, "7", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"7", "8", "9"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"7": "g", "8": "h", "9": "i"}) })
	l.Dump()

	l.Remove("7") // remove from head: 7,8,9 => 8,9
	assert.Equal(t, "9", l.tail.key)
	assert.Equal(t, "8", l.head.key)
	assert.Condition(t, func() bool { return CompareLL(l.head, []string{"8", "9"}) })
	assert.Condition(t, func() bool { return CompareMapValues(l.store, map[string]string{"8": "h", "9": "i"}) })
	l.Dump()
}

// CompareLL can follow a given linked list and compare values in the list against keys provided
func CompareLL(llHead *Entry, keys []string) bool {
	i := 0
	for n := llHead; n != nil; n = n.next {
		if i >= len(keys) {
			return false
		}
		if n.key != keys[i] {
			return false
		}
		i++
	}
	if i != len(keys) {
		return false
	}
	return true
}

// CompareMapValues compares a given Entry map with another map whose values just contain strings
func CompareMapValues(a map[string]*Entry, e map[string]string) bool {
	if len(a) != len(e) {
		return false
	}
	for k, v := range a {
		ev, exists := e[k]
		if !exists {
			return false
		}
		if ev != v.value {
			return false
		}
	}
	return true
}
