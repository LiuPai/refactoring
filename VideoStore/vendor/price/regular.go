package price

type regularPrice struct {
	basePrice
}

func (c *regularPrice) GetPriceCode() int {
	return Regular
}

func (c *regularPrice) GetCharge(daysRented int) float64 {
	result := float64(2)
	if daysRented > 2 {
		result += float64(daysRented-2) * 1.5
	}
	return result
}
