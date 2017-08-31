package unionfind

type QuickFindUF struct {
	id []int
}

func (qf *QuickFindUF) Connected(p, q int) bool {
	return qf.id[p] == qf.id[q]
}

func (qf *QuickFindUF) Union(p, q int) {
	idp := qf.id[p]
	idq := qf.id[q]
	for i, _ := range qf.id {
		if qf.id[i] == idp {
			qf.id[i] = idq
		}
	}
}

func NewQuickFindUF(n int) *QuickFindUF {
	var id []int
	for i := 0; i < n; i++ {
		id = append(id, i)
	}
	return &QuickFindUF{id: id}
}
