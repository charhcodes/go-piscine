package main

import (
	"fmt"
	"os"
)

/*
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
}*/

func main() {
	stringTargets := []string{"01", "galaxy", "galaxy 01"}
	alert := false

	for i := 1; i < len(os.Args); i++ {
		for j := 0; j < 3; j++ { // 3 represents number of stringtargets ie. 2 = galaxy
			if os.Args[i] == stringTargets[j] {
				alert = true
			}
		}
	}
	if alert {
		fmt.Println("Alert!!!")
	}
}
