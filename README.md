HTTP GET Request Example
This is a simple example program that demonstrates how to send an HTTP GET request to a URL and log the response body to the console using a custom logWriter type.

Usage
To use this program, simply run the main function in the main.go file. The program will send a GET request to http://google.com and log the response body to the console using the logWriter type.

Code Explanation
The main function sends a GET request to http://google.com using the http.Get function. If there was an error sending the request, the error is printed to the console and the program exits. If there was no error, the response body is copied to a logWriter instance using the io.Copy function. The logWriter instance logs the contents of the response body to the console.

The logWriter type defines a Write method that takes a byte slice as an argument and logs the contents of the byte slice to the console using the fmt.Println function. The Write method also returns the number of bytes written and no error to satisfy the io.Writer interface.

Dependencies
This program depends on the following packages:

fmt: for printing to the console
io: for copying data between sources and destinations
net/http: for sending HTTP requests
os: for exiting the program

License
This program is licensed under the MIT License. See the LICENSE file for more information.