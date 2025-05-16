package shipmentcalc_test

import (
	"slices"
	"testing"

	"shipment-calculator/shipmentcalc"
)

func TestBFSMinExcess(t *testing.T) {
	commonNums := []int{250, 500, 1000, 2000, 5000}
	testCases := map[string]struct {
		nums   []int
		target int
		want   *shipmentcalc.State
	}{
		"target=1": {
			nums:   commonNums,
			target: 1,
			want: &shipmentcalc.State{
				Total: 250,
				Count: 1,
				Path:  []int{250},
			},
		},
		"target=250": {
			nums:   commonNums,
			target: 250,
			want: &shipmentcalc.State{
				Total: 250,
				Count: 1,
				Path:  []int{250},
			},
		},
		"target=251": {
			nums:   commonNums,
			target: 251,
			want: &shipmentcalc.State{
				Total: 500,
				Count: 1,
				Path:  []int{500},
			},
		},
		"target=501": {
			nums:   commonNums,
			target: 501,
			want: &shipmentcalc.State{
				Total: 750,
				Count: 2,
				Path:  []int{250, 500},
			},
		},
		"target=12001": {
			nums:   commonNums,
			target: 12001,
			want: &shipmentcalc.State{
				Total: 12250,
				Count: 4,
				Path:  []int{250, 2000, 5000, 5000},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := shipmentcalc.BFSMinExcess(tc.nums, tc.target)
			if tc.want.Total != got.Total || tc.want.Count != got.Count || !slices.Equal(tc.want.Path, got.Path) {
				t.Errorf("want: %+v, but got: %+v", tc.want, got)
			}
		})
	}
}
