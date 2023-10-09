package main

import "fmt"

func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("zero div error")
	}
	return a / b, nil
}

//Templating
func area[T int | float32 | float64](breath T, width T) T {
	return breath * width
}

func main() {
	res, err := div(10, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	fmt.Println(area(10, 5))
	fmt.Printf("%0.2f\n", area(10.8, 5.4))
}
