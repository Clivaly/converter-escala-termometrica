package main

import (
	"fmt"
	"strings"
)

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func celsiusToKelvin(c float64) float64 {
	return c + 273.15
}

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func fahrenheitToKelvin(f float64) float64 {
	return (f - 32) * 5 / 9 + 273.15
}

func kelvinToFahrenheit(k float64) float64 {
	return (k-273.15)*9/5 + 32
}

func main() {
	var temperature float64
	var scale string

	fmt.Println("Enter temperature:")
	fmt.Scan(&temperature)

	fmt.Println("Enter scale (C, F, K):")
	fmt.Scan(&scale)

	switch strings.ToUpper(scale)  {
	case "C":
		fmt.Printf("%.2f°C is %.2f°F\n", temperature, celsiusToFahrenheit(temperature))
		fmt.Printf("%.2f°C is %.2fK\n", temperature, celsiusToKelvin(temperature))
	case "F":
		fmt.Printf("%.2f°F is %.2f°C\n", temperature, fahrenheitToCelsius(temperature))
		fmt.Printf("%.2f°F is %.2fK\n", temperature, fahrenheitToKelvin(temperature))
	case "K":
		fmt.Printf("%.2fK is %.2f°C\n", temperature, kelvinToCelsius(temperature))
		fmt.Printf("%.2fK is %.2f°F\n", temperature, kelvinToFahrenheit(temperature))
	default:
		fmt.Println("Invalid scale")
	}
}
