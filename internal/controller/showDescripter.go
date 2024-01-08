package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func searchMovies(query string) error {
	apiKey := "YOUR_API_KEY" // Replace with your TMDB API key

	// Build the URL for the search request
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s&query=%s", apiKey, query)

	// Send the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Process the response
	fmt.Println(string(body)) // Replace with your desired logic

	return nil
}

func main() {
	query := "Avengers" // Replace with your desired movie query
	err := searchMovies(query)
	if err != nil {
		fmt.Println("Error:", err)
	}
}