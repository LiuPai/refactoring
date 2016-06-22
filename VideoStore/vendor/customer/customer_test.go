package customer

import (
	"encoding/base64"
	"math/rand"
	"movie"
	"rental"
	"testing"
)

var (
	customerName = []string{"Jim Green", "Lily King", "Lucy King", "Li Lei"}
	movieTitle   = []string{
		// new release
		"Gleason", "All Eyez on Me", "Westworld", "Finding Dory",
		// regular
		"The Godfather", "The Dark Knight", "The Shawshank Redemption",
		// childrens
		"The Lion King", "WALL·E", "Up",
	}
	movieCode = []int{
		1, 1, 1, 1,
		0, 0, 0,
		2, 2, 2,
	}
	movies    []*movie.Movie
	customers []*Customer
)

type statementTest struct {
	c         *Customer
	statement string
}

var (
	// TestStatement
	statementTests []*statementTest
)

func loadMovies() []*movie.Movie {
	movies := make([]*movie.Movie, 0, 10)
	for i, title := range movieTitle {
		movies = append(movies, movie.NewMovie(title, movieCode[i]))
	}
	return movies
}

func createCustomers() []*Customer {
	customers := make([]*Customer, 0, 5)
	for _, name := range customerName {
		c := NewCustomer(name)
		chooseCount := rand.Intn(len(movies))
		indices := rand.Perm(len(movies))
		for i := 0; i < chooseCount; i++ {
			c.AddRental(rental.NewRental(
				movies[indices[i]], rand.Intn(5)+1))
		}
		customers = append(customers, c)
	}
	return customers
}

func init() {
	// set seed
	rand.Seed(0)
	// load movie
	movies = loadMovies()
	// create customers
	customers = createCustomers()
	statementTests = []*statementTest{
		{customers[0], "UmVudGFsIFJlY29yZCBmb3IgSmltIEdyZWVuCglUaGUgTGlvbiBLaW5nCTEuNQoJVGhlIERhcmsgS25pZ2h0CTYuNQoJQWxsIEV5ZXogb24gTWUJMTIKCVdBTEzCt0UJMS41CkFtb3VudCBvd2VkIGlzIDIxLjUKWW91IGVhcm5lZCA1IGZyZXF1ZW50cmVudGVycG9pbnRzIHJlbnRlciBwb2ludHM="},
		{customers[1], "UmVudGFsIFJlY29yZCBmb3IgTGlseSBLaW5nCglUaGUgRGFyayBLbmlnaHQJMy41CglXZXN0d29ybGQJMwoJV0FMTMK3RQkxLjUKCUdsZWFzb24JOQoJVGhlIExpb24gS2luZwkxLjUKCVRoZSBHb2RmYXRoZXIJNQpBbW91bnQgb3dlZCBpcyAyMy41CllvdSBlYXJuZWQgNyBmcmVxdWVudHJlbnRlcnBvaW50cyByZW50ZXIgcG9pbnRz"},
		{customers[2], "UmVudGFsIFJlY29yZCBmb3IgTHVjeSBLaW5nCglUaGUgU2hhd3NoYW5rIFJlZGVtcHRpb24JNQoJV2VzdHdvcmxkCTE1CglHbGVhc29uCTkKCVdBTEzCt0UJMwoJRmluZGluZyBEb3J5CTMKQW1vdW50IG93ZWQgaXMgMzUKWW91IGVhcm5lZCA3IGZyZXF1ZW50cmVudGVycG9pbnRzIHJlbnRlciBwb2ludHM="},
		{customers[3], "UmVudGFsIFJlY29yZCBmb3IgTGkgTGVpCglBbGwgRXlleiBvbiBNZQkxNQoJVGhlIExpb24gS2luZwkxLjUKQW1vdW50IG93ZWQgaXMgMTYuNQpZb3UgZWFybmVkIDMgZnJlcXVlbnRyZW50ZXJwb2ludHMgcmVudGVyIHBvaW50cw=="},
	}

}

func TestStatement(t *testing.T) {
	for i, test := range statementTests {
		v := test.c.Statement()
		eData := base64.StdEncoding.EncodeToString([]byte(v))
		if eData != test.statement {
			t.Errorf("Statement Test[%d] expected %s got %s.",
				i, test.statement, eData)
		}
	}
}
