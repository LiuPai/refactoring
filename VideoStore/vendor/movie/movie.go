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
