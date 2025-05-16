package shipmentcalc

import (
	"container/heap"
	"slices"
)

type State struct {
	Total int // total sum
	Count int // number of elements used
	Path  []int
}

// To find combinations using the minimum number of elements,
// BFS is an good strategy as it explores all combinations at a given level.
// This guarantees the use of the fewest elements.
// Additionally, BFS with a priority queue prioritizes combinations
// that minimize the excess.
func BFSMinExcess(nums []int, target int) *State {
	priorityQueue := &PriorityQueue{}
	heap.Init(priorityQueue)
	heap.Push(priorityQueue, &State{})

	visited := make(map[int]int)

	var current *State
	for priorityQueue.Len() > 0 {
		current = heap.Pop(priorityQueue).(*State)

		// The priority queue ensures that the first state reaching or exceeding
		// the target is the best one
		if current.Total >= target {
			break
		}

		for _, num := range nums {
			newTotal := current.Total + num
			newCount := current.Count + 1

			// Skip this path if we've already reached the same total using
			// fewer or equal elements.
			if visitedCount, ok := visited[newTotal]; ok && visitedCount <= newCount {
				continue
			}
			visited[newTotal] = newCount
			newPath := slices.Clone(current.Path)
			newPath = append(newPath, num)

			heap.Push(priorityQueue, &State{
				Total: newTotal,
				Count: newCount,
				Path:  newPath,
			})
		}
	}

	return current
}
