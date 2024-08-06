package main

import (
	"fmt"
)

func tempcTempf(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func tempcTempk(celsius float64) float64 {
	return celsius + 273.15
}

func tempfTempc(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func temfTempk(fahrenheit float64) float64 {
	return (fahrenheit-32)*5/9 + 273.15
}

func main() {
	var temp float64
	var scale string

	fmt.Print("coloque a temperatura: ")
	fmt.Scanln(&temp)
	fmt.Print("entre com a escala (C/F)")
	fmt.Scanln(&scale)

	if scale == "C" {
		fmt.Printf("graus celsius: %.2f\n", temp)
		fmt.Printf("graus Fahrenheit: %.2f\n", tempcTempf(temp))
		fmt.Printf("graus Kelvin: %.2f\n", tempcTempk(temp))
	} else if scale == "F" {
		fmt.Printf("graus Fahrenheit: %.2f\n", temp)
		fmt.Printf("graus celsius: %.2f\n", tempfTempc(temp))
		fmt.Printf("graus Kelvin: %.2f\n", temfTempk(temp))
	} else {
		fmt.Println("escala invalida")
	}
}
