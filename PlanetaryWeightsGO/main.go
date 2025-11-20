package main

import (
	"fmt"
)

func main() {
	//these are constants DO NOT CHANGE
	planetmap := map[string]float32{
		"MERCURY": 0.38,
		"VENUS":   0.91,
		"MOON":    0.165,
		"MARS":    0.38,
		"JUPITER": 2.34,
		"Saturn":  0.93,
		"URANUS":  0.92,
		"NEPTUNE": 1.12,
		"PLUTO":   0.666,
	}
	//Input to get name and weight values
	var Name string
	fmt.Print("What is your name? ")
	fmt.Scanln(&Name)

	var weight float32
	fmt.Print("what is your weight: ")
	fmt.Scanln(&weight)

	//iterate through map and print weight on diff planets
	for key, value := range planetmap {
		fmt.Printf("your weight on %s is %.2f \n", key, value*weight)
	}

}
