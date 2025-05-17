package shipmentcalc

// Having a struct makes it easy to add methods for other calculations,
// such as costs, time, etc.
type ShipmentCalculator struct {
	// We can move `packSizes` here in case the values don’t change frequently.
}

func New() *ShipmentCalculator {
	return &ShipmentCalculator{}
}

func (s *ShipmentCalculator) CalculateTotalPacks(packSizes []int, itemsOrder int) map[int]int {
	// At its core, the problem is to find the combination of numbers in 'packSizes'
	// that sums to at least 'itemsOrder', with the following priorities (in order):
	//
	// 1. Minimize the amount by which the total exceeds 'itemsOrder'.
	// 2. Minimize the number of elements used.
	state := BFSMinExcess(packSizes, itemsOrder)

	return countFrequency(state.Path)
}

func countFrequency(nums []int) map[int]int {
	frequencies := make(map[int]int)
	for _, val := range nums {
		frequencies[val]++
	}

	return frequencies
}
