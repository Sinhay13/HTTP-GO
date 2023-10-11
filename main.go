package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// Send a GET request to google.com
	resp, err := http.Get("http://google.com")

	// If there was an error, print it and exit the program
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Print the response object
	//fmt.Println(resp)

	// bs := make([]byte, 99999) // Create a byte slice with a capacity of 99999 bytes
	// resp.Body.Read(bs)        // Read the response body into the byte slice
	// fmt.Println(string(bs))   // Convert the byte slice to a string and print it to the console

	// Other way to do the same with ip package
	//io.Copy(os.Stdout, resp.Body) // Copy the response body to the console output

	lw := logWriter{}      // Create a new instance of the logWriter type
	io.Copy(lw, resp.Body) // Copy the response body to the logWriter instance
}

func (logWriter) Write(bs []byte) (int, error) {
	// Convert the byte slice to a string and print it to the console
	fmt.Println(string(bs))

	// Print the number of bytes written to the console
	fmt.Println("Just wrote this many bytes: ", len(bs))

	// Return the number of bytes written and no error
	return len(bs), nil
}
