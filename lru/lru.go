package lru

import (
	"errors"
	"time"
)

type AccessNode struct {
	prev *AccessNode
	next *AccessNode
	time time.Time
}

type LRU struct {
	data     map[string]string
	access   map[string]*AccessNode
	head     *AccessNode
	size     int
	capacity int
}

func NewLRU(capacity int) *LRU {
	lru := &LRU{
		data:   make(map[string]string),
		access: make(map[string]*AccessNode),
	}
	lru.capacity = capacity
	lru.size = 0
	lru.head = nil

	return lru
}

// promote this access node to the head of access links.
func (lru *LRU) promote(k string, anode *AccessNode) {
	head := lru.head

	anode.time = time.Now()
	anode.next = head
	anode.prev = nil

	if head != nil {
		head.prev = anode
	}

	lru.head = anode

	lru.access[k] = anode
}

// helper function. reverse access map, from k:v => v:k
func (lru LRU) reverseAccessMap() map[*AccessNode]string {
	reversed := make(map[*AccessNode]string)

	for k, anode := range lru.access {
		reversed[anode] = k
	}
	return reversed
}

// if the lru queue is too long, remove those least recent used keys
func (lru *LRU) clean() {
	head := lru.head

	size := 1
	for head != nil && size <= lru.capacity {
		head = head.next
		size++
	}

	if head != nil {
		reversed := lru.reverseAccessMap()

		for head != nil {
			head.prev.next = nil
			head.prev = nil

			k := reversed[head]
			delete(lru.access, k)
			delete(lru.data, k)

			head = head.next
		}
	}

	lru.size = size - 1
}

// put k:v
func (lru *LRU) Put(k string, v string) {
	anode, ok := lru.access[k]

	if ok {
		anode.prev.next = anode.next
		anode.next.prev = anode.prev
	} else {
		anode = &AccessNode{}
		lru.size++
	}

	lru.promote(k, anode)

	lru.data[k] = v

	if lru.size > lru.capacity {
		lru.clean()
	}
}

// get k
func (lru *LRU) Get(k string) (string, error) {
	anode, ok := lru.access[k]

	if ok {
		anode.prev.next = anode.next
		anode.next.prev = anode.prev

		lru.promote(k, anode)
		return lru.data[k], nil

	} else {
		return "", errors.New("not found")
	}
}

// get first key
func (lru *LRU) First() (string, error) {
	reversed := lru.reverseAccessMap()

	head := lru.head
	if head == nil {
		return "", errors.New("empty lru")
	}

	k := reversed[head]
	return lru.data[k], nil
}

// list all keys, ordered by last recent used.
func (lru LRU) Keys() []string {
	reversed := lru.reverseAccessMap()

	var key string
	var keys []string

	head := lru.head
	for head != nil {
		key = reversed[head]
		keys = append(keys, key)

		head = head.next
	}

	return keys
}
