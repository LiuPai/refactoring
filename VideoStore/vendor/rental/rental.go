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
	return r.movie.GetCharge(r.daysRented)
}

// GetFrequentRenterPoints return frequentrenterpoints this rental got
func (r *Rental) GetFrequentRenterPoints() int {
	return r.movie.GetFrequentRenterPoints(r.daysRented)
}
