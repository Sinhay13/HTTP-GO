package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Send a GET request to google.com
	resp, err := http.Get("http://google.com")

	// If there was an error, print it and exit the program
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Print the response object
	fmt.Println(resp)
}
