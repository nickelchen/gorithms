package unionfind

import (
	"log"
	"testing"
)

func TestQuickFind(_ *testing.T) {
	qf := NewQuickFindUF(10)
	qf.Union(4, 9)
	qf.Union(1, 7)
	qf.Union(2, 3)
	qf.Union(4, 3)
	// array is now: [0 7 3 3 3 5 6 7 8 3]

	if !qf.Connected(3, 9) {
		log.Fatal("QuickFind: ", qf.id, ", 3 & 9 should connected!")
	}

	if qf.Connected(1, 9) {
		log.Fatal("QuickFind: ", qf.id, ", 1 & 9 not connected!")
	}

}
