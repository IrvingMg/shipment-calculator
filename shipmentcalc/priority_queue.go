package shipmentcalc

type PriorityQueue []*State

func (pq PriorityQueue) Len() int {
	return len(pq)
}

// This is a core part of finding the best solution, as it defines the priorities.
// Priorities: lower `Total` is better, if equal, lower 'Count'.
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Total != pq[j].Total {
		return pq[i].Total < pq[j].Total
	}

	return pq[i].Count < pq[j].Count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // helps to release memory
	*pq = old[:n-1]

	return item
}
