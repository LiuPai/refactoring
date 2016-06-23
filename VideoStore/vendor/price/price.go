package price

import (
	"errors"
)

// move categories
const (
	Regular = iota
	NewRelease
	Childrens
)

// Price is a state pattern of movie
type Price interface {
	// GetPriceCode return current price's pricecode
	GetPriceCode() int
	// GetCharge return current price renten for daysRented need charge
	GetCharge(daysRented int) float64
	// GetFrequentRenterPoints return renterPoints should own for daysRented
	GetFrequentRenterPoints(daysRented int) int
}

type basePrice struct{}

func (b *basePrice) GetFrequentRenterPoints(daysRented int) int {
	return 1
}

// GetPrice generate price for select pricecode
func GetPrice(pricecode int) (Price, error) {
	switch pricecode {
	case Regular:
		return new(regularPrice), nil
	case Childrens:
		return new(childrensPrice), nil
	case NewRelease:
		return new(newReleasePrice), nil
	default:
		return nil, errors.New("Incorrect Price Code")
	}
}
