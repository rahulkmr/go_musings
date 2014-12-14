package main

import "fmt"

func main() {
	els := map[string]map[string]string{
		"H":  {"name": "Hydrogen", "state": "gas"},
		"He": {"name": "Helium", "state": "gas"},
		"Li": {"name": "Lithium", "state": "solid"},
	}
	fmt.Println(els["Li"])
	if val, ok := els["Un"]; ok {
		fmt.Println(val)
	}
}
