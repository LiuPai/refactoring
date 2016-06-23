package price

type childrensPrice struct {
	basePrice
}

func (c *childrensPrice) GetPriceCode() int {
	return Childrens
}

func (c *childrensPrice) GetCharge(daysRented int) float64 {
	result := 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * 1.5
	}
	return result
}
