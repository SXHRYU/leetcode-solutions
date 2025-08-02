package structures

import (
	"container/heap"
)

type entry struct {
	key      string
	priority int

	index int
}

type PQ []*entry

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	temp := x.(*entry)
	temp.index = len(*pq)
	*pq = append(*pq, temp)
}

func (pq *PQ) Pop() any {
	temp := (*pq)[len(*pq)-1]
	temp.index = -1
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

func (pq *PQ) update(entry *entry, value string, priority int) {
	entry.key = value
	entry.priority = priority
	heap.Fix(pq, entry.index)
}
