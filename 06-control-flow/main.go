package main

import "fmt"

func main() {
	var (
		age    int     = 20
		weight float64 = 22.5
	)
	if age > 18 {
		fmt.Println("You are a major")
	} else {
		fmt.Println("You are a minor")
	}
	switch {
	case weight < 3.0:
		fmt.Println("You are too feeble")
		fallthrough
	case weight < 5.0:
		fmt.Println("You are underweight")
	case weight > 5.0 && weight < 50:
		fmt.Println("You are having the right weight")
	case weight == 0.0:
		fmt.Println("Invalid weight")
	default:
		fmt.Println("Enter weight below 50")
	}
	/*
		for {
			fmt.Println("Never ending loop")
		}
	*/

	// white True in python while 1 in C
	for c := 0; c < 10; c++ {
		fmt.Println("Counted loop")
	}
	myarr := [5]int{2, 3, 4, 5, 6}
	myslice := []string{"Kri", "Siva"}
	mymap := map[string]string{"Name": "Jega", "Age": "20"}

	for idx, value := range myarr {
		fmt.Println(idx, value)
	}
	for _, value := range myslice {
		fmt.Println(value)
	}
	for key, value := range mymap {
		fmt.Println(key, value)
	}

}
