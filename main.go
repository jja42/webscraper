package main

import (
	"fmt"
	"os"
)

func main() {
	//get command line arguments
	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(arguments) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s\n", arguments[1])

	output, err := getHTML(arguments[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Output:\n %s", output)
}
