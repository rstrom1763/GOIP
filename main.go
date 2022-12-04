package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// This is just a placeholder for now
func main() {
	args := os.Args
	uri := args[1]

	if len(args) < 2 || len(args) > 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}

	//Validate whether url is valid or not
	if _, err := url.ParseRequestURI(uri); err != nil {
		fmt.Println("URL is not valid! ")
		os.Exit(1)
	}

	response, err := http.Get(uri)
	if err != nil {
		fmt.Println("Request Failed!")
		os.Exit(1)
	} else {
		fmt.Println(response)
	}
}
