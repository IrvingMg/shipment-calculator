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
	return map[int]int{0: 0}
}
