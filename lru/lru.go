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

func (lru *LRU) updateAccess(k string, anode *AccessNode) {
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

func (lru *LRU) clean() {
	head := lru.head

	size := 1
	for head != nil && size <= lru.capacity {
		head = head.next
		size++
	}

	reversed := make(map[*AccessNode]string)

	for k, anode := range lru.access {
		reversed[anode] = k
	}

	if head != nil {
		head.prev.next = nil
		head.prev = nil

		// for head != nil {
		// 	k := reversed[head]
		// 	delete(lru.access, k)
		// 	head = head.next
		// }

		head = nil
	}

	lru.size = size - 1
}

func (lru *LRU) Put(k string, v string) {
	anode, ok := lru.access[k]

	if ok {
		anode.prev.next = anode.next
		anode.next.prev = anode.prev
	} else {
		anode = &AccessNode{}
		lru.size++
	}

	lru.updateAccess(k, anode)

	lru.data[k] = v

	if lru.size > lru.capacity {
		lru.clean()
	}
}

func (lru *LRU) Get(k string) (string, error) {
	anode, ok := lru.access[k]
	if !ok {
		return "", errors.New("not found")
	}

	lru.updateAccess(k, anode)

	return lru.data[k], nil
}

func (lru *LRU) First() (string, error) {
	head := lru.head
	for k, v := range lru.access {
		if v == head {
			return lru.data[k], nil
		}
	}
	return "", errors.New("empty lru")
}

func (lru LRU) Keys() []string {
	reversed := make(map[*AccessNode]string)

	for k, anode := range lru.access {
		reversed[anode] = k
	}

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
