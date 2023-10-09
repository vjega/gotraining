package main

import "fmt"

func main() {

	//Simple date type on stack
	var (
		name          string  = "Kri"
		age           int     = 19
		weight        float32 = 52.5
		marriedstatus bool    = false
	)
	// Collectons
	var (
		myarr   [10]int           = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		myslice []int             = []int{0, 1}
		mymap   map[string]string = map[string]string{
			"Name":  "Jega",
			"Title": "Software Eng",
		}
	)
	fmt.Println(name, age, weight, marriedstatus)
	fmt.Println(myarr, myslice, mymap)
}
