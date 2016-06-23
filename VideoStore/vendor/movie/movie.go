package movie

import (
	"price"
)

// Movie store base info and provide price methods
type Movie struct {
	title      string
	priceState price.Price
}

// NewMovie create movie with title and priceCode.
func NewMovie(title string, priceCode int) *Movie {
	m := &Movie{
		title: title,
	}
	m.SetPriceCode(priceCode)
	return m
}

// SetPriceCode set movie's price state
func (m *Movie) SetPriceCode(pricecode int) {
	p, err := price.GetPrice(pricecode)
	if err != nil {
		panic(err)
	}
	m.priceState = p
}

// GetPriceCode return movie's price state
func (m *Movie) GetPriceCode() int {
	return m.priceState.GetPriceCode()
}

// Title return title of movie
func (m *Movie) Title() string {
	return m.title
}

// GetCharge return price of this rental with daysRented defined
func (m *Movie) GetCharge(daysRented int) float64 {
	return m.priceState.GetCharge(daysRented)
}

// GetFrequentRenterPoints return frequentrenterpoints this rental got with
// daysRented defined
func (m *Movie) GetFrequentRenterPoints(daysRented int) int {
	return m.priceState.GetFrequentRenterPoints(daysRented)
}
