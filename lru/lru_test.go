package lru

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {

	l := NewLRU(10)
	l.Put("abc1", "efg")
	l.Put("abc2", "efg")
	l.Put("abc3", "efg")
	l.Put("abc4", "efg")
	l.Put("abc5", "efg")
	l.Put("abc6", "efg")
	l.Put("abc7", "efg")
	l.Put("abc8", "efg")
	l.Put("abc9", "efg")
	l.Put("abc10", "efg")

	l.Get("abc8")
	l.Get("abc4")
	l.Get("abc2")

	fmt.Printf("lru keys :%v", l.Keys())
}
