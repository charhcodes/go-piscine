package piscine

import (
	"fmt"
	"os"
)

func comcheck() {
	arguments := os.Args
	stringSlice := []string(arguments)

	for i := 0; i < len(arguments); i++ {
		if stringSlice[i] == "01" {
			fmt.Println("Alert!!!")
			return
		} else if stringSlice[i] == "g" && stringSlice[i+1] == "a" && stringSlice[i+2] == "l" && stringSlice[i+3] == "a" && stringSlice[i+4] == "x" && stringSlice[i] == "y" {
			fmt.Println("Alert!!!")
			return
		} else {
			return
		}
	}
}
