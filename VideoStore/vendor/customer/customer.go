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
	result := "Rental Record for " + c.Name() + "\n"
	for _, each := range c.rentals {
		// show figures for this rental
		result += "\t" + each.Movie().Title() + "\t" +
			strconv.FormatFloat(each.GetCharge(), 'f', -1, 64) + "\n"
	}
	// add footer lines
	result += "Amount owed is " +
		strconv.FormatFloat(c.getTotalCharge(), 'f', -1, 64) + "\n"
	result += "You earned " + strconv.Itoa(c.getTotalFrequentRenterPoints()) +
		" frequentrenterpoints renter points"
	return result
}

// HTMLStatement return HTML style statement of customer order
func (c *Customer) HTMLStatement() string {
	result := "<H1>Rentals for <EM>" + c.Name() + "</EM></H1><P>"
	for _, each := range c.rentals {
		// show figures for this rental
		result += each.Movie().Title() + ": " +
			strconv.FormatFloat(each.GetCharge(), 'f', -1, 64) + "<BR>"
	}
	// add footer lines
	result += "<P>You owe <EM>" +
		strconv.FormatFloat(c.getTotalCharge(), 'f', -1, 64) +
		"</EM></P>"
	result += "On this rental you earned <EM>" +
		strconv.Itoa(c.getTotalFrequentRenterPoints()) +
		"</EM> frequent renter points</P>"
	return result
}

func (c *Customer) getTotalCharge() float64 {
	var result float64
	for _, each := range c.rentals {
		result += each.GetCharge()
	}
	return result
}

func (c *Customer) getTotalFrequentRenterPoints() int {
	var result int
	for _, each := range c.rentals {
		result += each.GetFrequentRenterPoints()
	}
	return result
}
