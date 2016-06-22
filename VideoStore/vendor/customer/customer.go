package customer

import (
	"strconv"

	"movie"
	"rental"
)

// Customer store customer information
type Customer struct {
	name    string
	rentals []*rental.Rental
}

// NewCustomer create new customer with name
func NewCustomer(name string) *Customer {
	return &Customer{
		name:    name,
		rentals: make([]*rental.Rental, 0, 3),
	}
}

// AddRental adds rental r to customer rentals
func (c *Customer) AddRental(r *rental.Rental) {
	c.rentals = append(c.rentals, r)
}

// Name return customer's name
func (c *Customer) Name() string {
	return c.name
}

// Statement return string statement of customer order
func (c *Customer) Statement() string {
	var (
		totalAmount          float64
		frequentRenterPoints int
	)
	result := "Rental Record for " + c.Name() + "\n"
	for _, each := range c.rentals {
		var (
			thisAmount float64
		)
		switch each.Movie().PriceCode {
		case movie.Regalur:
			thisAmount += 2
			if each.DaysRented() > 2 {
				thisAmount += float64(each.DaysRented()-2) * 1.5
			}
		case movie.NewRelease:
			thisAmount += float64(each.DaysRented() * 3)
		case movie.Childrens:
			thisAmount += 1.5
			if each.DaysRented() > 3 {
				thisAmount += float64(each.DaysRented()-3) * 1.5
			}
		}

		// add frequentrenterpoints renter points
		frequentRenterPoints++
		// add bonus for a two day new release rental
		if each.Movie().PriceCode == movie.NewRelease &&
			each.DaysRented() > 1 {
			frequentRenterPoints++
		}

		// show figures for this rental
		result += "\t" + each.Movie().Title() + "\t" +
			strconv.FormatFloat(thisAmount, 'f', -1, 64) + "\n"
		totalAmount += thisAmount
	}
	// add footer lines
	result += "Amount owed is " +
		strconv.FormatFloat(totalAmount, 'f', -1, 64) + "\n"
	result += "You earned " + strconv.Itoa(frequentRenterPoints) +
		" frequentrenterpoints renter points"
	return result
}
