package shipmentcalc

// Having a struct makes it easy to add methods for other calculations,
// such as costs, time, etc.
type shipmentCalculator struct {
	// We can move `packSizes` here in case the values don’t change frequently.
}

func New() shipmentCalculator {
	return shipmentCalculator{}
}

func (s *shipmentCalculator) CalculateTotalPacks(packSizes []int, itemsOrder int) map[int]int {
	return nil
}
