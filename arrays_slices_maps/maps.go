package main

import "fmt"

func main() {
	x := make(map[string]int) // x is a map of string to ints,
	// the key type is represented inside [] and the value type outside
	x["key"] = 10
	fmt.Println(x["key"])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["B"])
	// go returns an empty value for non-existent keys
	fmt.Println(elements["Un"])
	//better way to check the same
	name, ok := elements["Un"]
	fmt.Println(name, ok)
	// typical idiomatic way to write the same in go
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}

	// another map example
	detailedElements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "Gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, exists := detailedElements["Li"]; exists {
		fmt.Println(el["name"], el["state"])
	}

}
