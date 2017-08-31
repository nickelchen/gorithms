package unionfind

type QuickUnionUF struct {
	id []int
}

func (qu *QuickUnionUF) Connected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}

func (qu *QuickUnionUF) Union(p, q int) {
	rootp := qu.root(p)
	rootq := qu.root(q)
	qu.id[rootp] = rootq
}

func (qu *QuickUnionUF) root(p int) int {
	for p != qu.id[p] {
		p = qu.id[p]
	}
	return p
}

func NewQuickUnionUF(n int) *QuickUnionUF {
	var id []int
	for i := 0; i < n; i++ {
		id = append(id, i)
	}
	return &QuickUnionUF{id: id}
}
