package main

import (
	"fmt"
)

func main() {
	var temperature float64
	fmt.Print("Enter temperature in degrees: ")
	fmt.Scan(&temperature)
	if temperature < 32 {
		fmt.Println("weather too cold wear a jacket")
	} else if temperature > 32 && temperature < 50 {
		fmt.Println("weather a bit cold wear a light jacket")
	} else {
		fmt.Println("weather warm dont wear a jacket")
	}
}
