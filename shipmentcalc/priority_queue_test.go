package shipmentcalc_test

import (
	"container/heap"
	"testing"

	"shipment-calculator/shipmentcalc"
)

func TestPriorityQueue(t *testing.T) {
	testCases := map[string]struct {
		states []*shipmentcalc.State
		want   *shipmentcalc.State
	}{
		"case 1 - lower count is better": {
			states: []*shipmentcalc.State{
				{Count: 1, Total: 250},
				{Count: 1, Total: 500},
			},
			want: &shipmentcalc.State{Count: 1, Total: 250},
		},
		"case 2 - lower count is better": {
			states: []*shipmentcalc.State{
				{Count: 1, Total: 500},
				{Count: 2, Total: 500},
			},
			want: &shipmentcalc.State{Count: 1, Total: 500},
		},
		"case 3 - lower count is better": {
			states: []*shipmentcalc.State{
				{Count: 2, Total: 750},
				{Count: 1, Total: 1000},
				{Count: 3, Total: 750},
			},
			want: &shipmentcalc.State{Count: 2, Total: 750},
		},
		"case 4 - lower total is better": {
			states: []*shipmentcalc.State{
				{Count: 4, Total: 7250},
				{Count: 3, Total: 15000},
			},
			want: &shipmentcalc.State{Count: 4, Total: 7250},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			priorityQueue := &shipmentcalc.PriorityQueue{}
			heap.Init(priorityQueue)
			for _, state := range tc.states {
				heap.Push(priorityQueue, state)
			}

			got := heap.Pop(priorityQueue).(*shipmentcalc.State)
			if tc.want.Count != got.Count || tc.want.Total != got.Total {
				t.Errorf("want: %+v, but got: %+v", tc.want, got)
			}
		})
	}
}
