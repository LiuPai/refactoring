package rental

import (
	"movie"
)

// Rental movie rental state
type Rental struct {
	movie      *movie.Movie
	daysRented int
}

// NewRental create a new Rental
func NewRental(m *movie.Movie, daysRented int) *Rental {
	return &Rental{
		movie:      m,
		daysRented: daysRented,
	}
}

// DaysRented return days the rental is
func (r *Rental) DaysRented() int {
	return r.daysRented
}

// Movie return retal movie
func (r *Rental) Movie() *movie.Movie {
	return r.movie
}

// GetCharge return price of this rental
func (r *Rental) GetCharge() float64 {
	var result float64
	// determine amounts for each line
	switch r.Movie().PriceCode {
	case movie.Regalur:
		result += 2
		if r.DaysRented() > 2 {
			result += float64(r.DaysRented()-2) * 1.5
		}
	case movie.NewRelease:
		result += float64(r.DaysRented() * 3)
	case movie.Childrens:
		result += 1.5
		if r.DaysRented() > 3 {
			result += float64(r.DaysRented()-3) * 1.5
		}
	}
	return result
}
