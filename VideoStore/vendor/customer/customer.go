package customer

import (
	"strconv"

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
		// show figures for this rental
		frequentRenterPoints += each.GetFrequentRenterPoints()
		result += "\t" + each.Movie().Title() + "\t" +
			strconv.FormatFloat(each.GetCharge(), 'f', -1, 64) + "\n"
		totalAmount += each.GetCharge()
	}
	// add footer lines
	result += "Amount owed is " +
		strconv.FormatFloat(totalAmount, 'f', -1, 64) + "\n"
	result += "You earned " + strconv.Itoa(frequentRenterPoints) +
		" frequentrenterpoints renter points"
	return result
}
