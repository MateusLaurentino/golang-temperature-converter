package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Informe a temperatura em ºC: ")
	scanner.Scan()
	input := scanner.Text()

	temperature, error := strconv.ParseFloat(input, 1000)
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	fmt.Println("Graus Celsius: ", temperature)
	fmt.Println("Em Fahrenheit:", convertInFahrenheit(temperature))
	fmt.Println("Em Kelvin:", convertInKelvin(temperature))

}

func convertInFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func convertInKelvin(celsius float64) float64 {
	return celsius + 273.15
}
