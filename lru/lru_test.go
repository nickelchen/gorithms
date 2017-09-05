package lru

import (
	"reflect"
	"testing"
)

func TestLRU(t *testing.T) {

	l := NewLRU(8)
	l.Put("k1", "value")
	l.Put("k2", "value")
	l.Put("k3", "value")
	l.Put("k4", "value")
	l.Put("k5", "value")
	l.Put("k6", "value")
	l.Put("k7", "value")
	l.Put("k8", "value")
	l.Put("k9", "value")
	l.Put("k10", "value")

	expect := []string{"k10", "k9", "k8", "k7", "k6", "k5", "k4", "k3"}
	got := l.Keys()
	if !reflect.DeepEqual(got, expect) {
		t.Fatal("not equal")
	}

	l.Get("k8")
	l.Get("k4")
	l.Get("k2")

	expect = []string{"k4", "k8", "k10", "k9", "k7", "k6", "k5", "k3"}
	got = l.Keys()
	if !reflect.DeepEqual(got, expect) {
		t.Fatal("not equal")
	}

	// this will be ignored.
	l.Get("k1")
	l.Put("k7", "value")

	expect = []string{"k7", "k4", "k8", "k10", "k9", "k6", "k5", "k3"}
	got = l.Keys()
	if !reflect.DeepEqual(got, expect) {
		t.Fatal("not equal")
	}

}
