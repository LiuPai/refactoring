package movie

// move categories
const (
	Regalur = iota
	NewRelease
	Childrens
)

// Movie store base info and provide price methods
type Movie struct {
	title     string
	PriceCode int
}

// NewMovie create movie with title and priceCode.
func NewMovie(title string, priceCode int) *Movie {
	return &Movie{
		title:     title,
		PriceCode: priceCode,
	}
}

// Title return title of movie
func (m *Movie) Title() string {
	return m.title
}

// GetCharge return price of this rental with daysRented defined
func (m *Movie) GetCharge(daysRented int) float64 {
	var result float64
	// determine amounts for each line
	switch m.PriceCode {
	case Regalur:
		result += 2
		if daysRented > 2 {
			result += float64(daysRented-2) * 1.5
		}
	case NewRelease:
		result += float64(daysRented * 3)
	case Childrens:
		result += 1.5
		if daysRented > 3 {
			result += float64(daysRented-3) * 1.5
		}
	}
	return result
}

// GetFrequentRenterPoints return frequentrenterpoints this rental got with
// daysRented defined
func (m *Movie) GetFrequentRenterPoints(daysRented int) int {
	if m.PriceCode == NewRelease && daysRented > 1 {
		return 2
	}
	return 1
}
