package main

import "fmt"
import "os"

func main() {
	// User input stored in variable's
	fmt.Println("Brandon's Temperature Converter")
	var UserTemp float32
	fmt.Print("Enter a temperature: ")
	fmt.Scanf("%f", &UserTemp)
	var FarCel string
	fmt.Println("Is the temperature F for Fahrenheit or C for Celcius?")
	fmt.Scanf("%s", &FarCel)

	// If/else statements with F/C Convertion
	if FarCel == "F" || FarCel == "f" {
		if UserTemp > 212.00 {
			fmt.Println("Temp can not be greater > 212")
		} else {
			fmt.Printf("The Celsius equivlant is: %.2f \n", (5.0/9)*(UserTemp-32))
		}
	} else if FarCel == "C" || FarCel == "C" {
		if UserTemp > 100.00 {
			fmt.Println("Temp can not be greater > 100")
		} else {
			fmt.Printf("The Fahrenheit equivalent is: %.2f \n", ((9.0/5.0)*UserTemp)+32)
		}
	} else {
		fmt.Println("Enter a F or C \nExiting Program")
		os.Exit(1)
	}

}
