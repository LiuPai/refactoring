package main

import (
	"math/rand"

	"customer"
	"movie"
	"rental"
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
	customers []*customer.Customer
)

func loadMovies() []*movie.Movie {
	movies := make([]*movie.Movie, 0, 10)
	for i, title := range movieTitle {
		movies = append(movies, movie.NewMovie(title, movieCode[i]))
	}
	return movies
}

func createCustomers() []*customer.Customer {
	customers := make([]*customer.Customer, 0, 5)
	for _, name := range customerName {
		c := customer.NewCustomer(name)
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
}

func main() {
	for _, c := range customers {
		println(c.Statement())
	}
}
