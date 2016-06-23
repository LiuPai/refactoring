package price

type newReleasePrice struct {
	basePrice
}

func (c *newReleasePrice) GetPriceCode() int {
	return NewRelease
}

func (c *newReleasePrice) GetCharge(daysRented int) float64 {
	return float64(daysRented) * 3
}

func (c *newReleasePrice) GetFrequentRenterPoints(daysRented int) int {
	if daysRented > 1 {
		return 2
	}
	return 1
}
