package shipmentcalc_test

import (
	"maps"
	"testing"

	"shipment-calculator/shipmentcalc"
)

func TestShipmentCalc(t *testing.T) {
	commonPackSizes := []int{250, 500, 1000, 2000, 5000}
	testCases := map[string]struct {
		packSizes  []int
		itemsOrder int
		want       map[int]int
	}{
		"amount=1": {
			packSizes:  commonPackSizes,
			itemsOrder: 1,
			want:       map[int]int{250: 1},
		},
		"amount=250": {
			packSizes:  commonPackSizes,
			itemsOrder: 250,
			want:       map[int]int{250: 1},
		},
		"amount=251": {
			packSizes:  commonPackSizes,
			itemsOrder: 251,
			want:       map[int]int{500: 1},
		},
		"amount=501": {
			packSizes:  commonPackSizes,
			itemsOrder: 501,
			want:       map[int]int{500: 1, 250: 1},
		},
		"amount=12001": {
			packSizes:  commonPackSizes,
			itemsOrder: 12001,
			want:       map[int]int{5000: 2, 2000: 1, 250: 1},
		},
		"amount=500000": {
			packSizes:  []int{23, 31, 53},
			itemsOrder: 500000,
			want:       map[int]int{23: 2, 31: 7, 53: 9429},
		},
	}

	calc := shipmentcalc.New()
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := calc.CalculateTotalPacks(tc.packSizes, tc.itemsOrder)
			if !maps.Equal(tc.want, got) {
				t.Errorf("want: %v, but got: %v", tc.want, got)
			}
		})
	}
}
