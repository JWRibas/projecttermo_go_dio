package main

import "fmt"

const (
	ebK float64 = 373.0
)

func main() {
	tempK := ebK
	tempC := tempK - 273.0

	fmt.Printf("A Temperatura de ebulição da agua em K = %g, já em C é = %g ", tempK, tempC)
}
