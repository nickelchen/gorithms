package unionfind

import (
	"log"
	"testing"
)

func TestQuickUnion(_ *testing.T) {
	uf := NewQuickUnionUF(10)
	uf.Union(4, 9)
	uf.Union(1, 7)
	uf.Union(2, 3)
	uf.Union(4, 3)
	// array is now: [0 7 3 3 3 5 6 7 8 3]

	if !uf.Connected(3, 9) {
		log.Fatal("QuickUnion: ", uf.id, ", 3 & 9 should connected!")
	}

	if uf.Connected(1, 9) {
		log.Fatal("QuickUnion: ", uf.id, ", 1 & 9 not connected!")
	}

}
