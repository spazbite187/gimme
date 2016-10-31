package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var version = "1.0.0-DEV"

func main() {
	fmt.Printf("gimme - version %s\n", version)

	if len(os.Args) < 2 {
		fmt.Println("url required")
		os.Exit(1)
	}

	url := os.Args[1]
	urlParts := strings.Split(url, "/")
	filename := urlParts[len(urlParts) - 1]

	fmt.Println("Downloading content from", url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error occured.\n", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Data saved to %s\n", filename)
}
