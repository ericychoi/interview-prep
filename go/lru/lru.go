package lru

import "fmt"

type Entry struct {
	key   string
	value string
	next  *Entry
	prev  *Entry
}

type LRU struct {
	limit int
	store map[string]*Entry
	head  *Entry // head of ll, first candidate to remove
	tail  *Entry // last Entry in ll, last candidate to remove
}

func (l *LRU) Add(k, v string) {
	e := Entry{value: v, key: k}
	l.store[k] = &e
	if len(l.store) > l.limit {
		// evict
		delete(l.store, l.head.key)
		l.head = l.head.next
		if l.head.prev != nil {
			l.head.prev = nil
		}
	}

	if len(l.store) == 1 {
		l.head = &e
		l.tail = &e
		return
	}

	l.attachToTail(&e)
}

// TODO: in order to make this thread safe, use mutex
func (l *LRU) Read(k string) string {
	e, ok := l.store[k]
	if !ok {
		return ""
	}

	l.moveToEnd(e)
	return e.value
}

// move the Entry in the LL to the end
func (l *LRU) moveToEnd(e *Entry) {
	// if I am at the end already, no need to do anything
	if l.tail == e {
		return
	}

	// take care of surrounding nodes
	if e.prev != nil {
		e.prev.next = e.next
	}
	e.next.prev = e.prev

	// take care of head and tail
	if l.head == e {
		l.head = e.next
	}
	l.attachToTail(e)
}

func (l *LRU) attachToTail(e *Entry) {
	// attach myself to the tail, tail is never nil
	e.prev = l.tail
	l.tail.next = e

	l.tail = e
	e.next = nil
}

func (l *LRU) Dump() {
	fmt.Printf("head: %+v, tail: %+v\n", l.head, l.tail)
	for n := l.head; n != nil; n = n.next {
		var pv, nv string
		if n.prev != nil {
			pv = n.prev.key
		}
		if n.next != nil {
			nv = n.next.key
		}
		fmt.Printf("%s(%s:%s) => ", n.key, pv, nv)
	}
	fmt.Println("")
}

func (l *LRU) Remove(k string) {
	e, ok := l.store[k]
	if !ok {
		return
	}
	delete(l.store, k)

	if len(l.store) == 0 {
		l.head = nil
		l.tail = nil
		return
	}

	// take care of surrounding nodes
	if e.prev != nil {
		e.prev.next = e.next
	}
	if e.next != nil {
		e.next.prev = e.prev
	}

	// head and tail
	if l.head == e {
		l.head = e.next
	}
	if l.tail == e {
		l.tail = e.prev
	}
}

func NewLRU(l int) *LRU {
	return &LRU{
		store: make(map[string]*Entry),
		limit: l,
	}
}
