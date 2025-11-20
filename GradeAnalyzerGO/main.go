package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Print("Name of the person we are calculating the grades for: ")
	fmt.Scan(&name)

	var Test1, Test2, Test3, Test4 float32
	fmt.Print("Test 1: ")
	fmt.Scan(&Test1)
	fmt.Print("Test 2: ")
	fmt.Scan(&Test2)
	fmt.Print("Test 3: ")
	fmt.Scan(&Test3)
	fmt.Print("Test 4: ")
	fmt.Scan(&Test4)

	var DropCon string
	fmt.Print("Do you wish to drop the lowest grade Y or N?  ")
	fmt.Scan(&DropCon)

	if DropCon != "Y" && DropCon != "N" {
		fmt.Println("Enter to drop the lowest grade Y or N.")
		os.Exit(1)
	}
	if Test1 <= 0 || Test2 <= 0 || Test3 <= 0 || Test4 <= 0 {
		fmt.Println("Test scores must be greater than 0.")
		os.Exit(1)
	}
	var smallest float32
	var Avg float32
	if DropCon == "Y" {
		if Test1 < Test2 {
			smallest = Test1
		} else {
			smallest = Test2
		}
		if Test3 < smallest {
			smallest = Test3
		}
		if Test4 < smallest {
			smallest = Test4
		}
		Avg = ((Test1 + Test2 + Test3 + Test4) - smallest) / 3
	} else {
		Avg = (Test1 + Test2 + Test3 + Test4) / 4
	}

	var Grade string
	if Avg >= 97.0 {
		Grade = "A+"
	} else if Avg >= 94.0 {
		Grade = "A"
	} else if Avg >= 90.0 {
		Grade = "A-"
	} else if Avg >= 87.0 {
		Grade = "B+"
	} else if Avg >= 84.0 {
		Grade = "B"
	} else if Avg >= 80.0 {
		Grade = "B-"
	} else if Avg >= 77.0 {
		Grade = "C+"
	} else if Avg >= 74.0 {
		Grade = "C"
	} else if Avg >= 70.0 {
		Grade = "C-"
	} else if Avg >= 67.0 {
		Grade = "D+"
	} else if Avg >= 64.0 {
		Grade = "D"
	} else if Avg >= 60.0 {
		Grade = "D-"
	} else {
		Grade = "F"
	}
	fmt.Printf("%s test average is %.2f\n", name, Avg)
	fmt.Printf("Letter Grade for the test is: %s\n", Grade)
}
