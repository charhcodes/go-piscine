package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) > 2 {
		fmt.Println("Too many arguments")
		return
	} else if len(arguments) == 1 {
		fmt.Println("File name missing")
		return
	}

	file, err := os.Open("quest8.txt") // opens files within os, cannot take multiple values usually
	if err != nil {
		fmt.Printf("")
	}

	// read arr, print result as a string
	arr := make([]byte, 14)
	file.Read(arr)           // reads arr, can only take bytes as arguments
	fmt.Println(string(arr)) // prints text in sample.txt

	// close the file, good practice esp when dealing w database performance
	file.Close()
}
