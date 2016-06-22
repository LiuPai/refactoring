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
		thisAmount := amountFor(each)

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

func amountFor(aRental *rental.Rental) float64 {
	return aRental.GetCharge()
}
