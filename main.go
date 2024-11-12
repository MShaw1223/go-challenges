package main

import (
	"fmt"
	"log"
	"os"

	datastructs "github.com/MShaw1223/go-challenges.git/data-structures"
	dbstorage "github.com/MShaw1223/go-challenges.git/db-storage"
	filesystem "github.com/MShaw1223/go-challenges.git/file-system"
	objects "github.com/MShaw1223/go-challenges.git/objects"
)

func main() {
	arg := os.Args[1]

	switch {
	case arg == "data-structures":
		fmt.Print("Datastructs module Starting...\n\n")
		datastructs.Array()
		datastructs.Tuple()
		datastructs.Hash_map()
		fmt.Print("\nDatastructs module complete...\n\n")
	case arg == "file-system":
		fmt.Print("Filesystem module Starting...\n\n")
		filesystem.FileSystem()
		fmt.Print("\nFilesystem module complete...\n\n")
	case arg == "objects":
		fmt.Print("Objects Modules Starting\n\n")
		fmt.Print("Enter a choice from the following: \n1) Bank\n2) Vehicle\n3) Shape\n4) Calculator\n5) Generics\n\n")
		var choice int
		fmt.Print("Choice from 1 - 5: ")
		fmt.Scanln(&choice)
		switch {
		case choice == 1:
			fmt.Print("Bank module Starting...\n\n")
			bank()
			fmt.Print("\nBank module complete...\n\n")
		case choice == 2:
			fmt.Print("Vehicle module Starting...\n\n")
			vehicle()
			fmt.Print("\nVehicle module complete...\n\n")
		case choice == 3:
			fmt.Print("Shapes module starting...\n\n")
			shape()
			fmt.Print("\nShapes module complete...\n\n")
		case choice == 4:
			fmt.Print("Calculator module starting...\n\n")
			calculator()
			fmt.Print("\nCalculator module complete...\n\n")
		case choice == 5:
			fmt.Print("Generics Module Starting...\n\n")
			Gen()
			fmt.Print("\nGenerics Module Complete...\n\n")
		default:
			fmt.Println("Wrong")
		}
		fmt.Print("\nObjects Modules Complete...\n\n")
	case arg == "db-storage":
		fmt.Print("Database Storage Module Starting...\n\n")
		if err := dbstorage.InitDB(); err != nil {
			log.Fatal("Failed to initialize database:", err)
		}
		dbstorage.CreateTable()
		fmt.Print("Database Storage Module Complete...\n\n")
	default:
		fmt.Print("Incorrect input\n\n")
		main()
	}
}

func Gen() {
	str_summary := objects.Container[objects.Str]{Val: "Hello"}.Summary()
	fmt.Println(str_summary)
	int_summary := objects.Container[objects.I]{Val: 68}.Summary()
	fmt.Println(int_summary)
	bool_summary := objects.Container[objects.Bl]{Val: false}.Summary()
	fmt.Println(bool_summary)
	fl32_summary := objects.Container[objects.Fl32]{Val: 3.3}.Summary()
	fmt.Println(fl32_summary)
	fl64_summary := objects.Container[objects.Fl64]{Val: 3.333333}.Summary()
	fmt.Println(fl64_summary)
}

func bank() {
	// make new person
	person := objects.BankAccount{
		AccountNum: 1,
		Balance:    200.32,
		InterestRt: 5.0,
	}

	// use methods in person
	person.See_Balance()

	person.Deposit(20000.68)
	person.See_Balance()

	person.Withdraw(15000)
	person.See_Balance()

	person.Returns(2000, 3.0)
}

func vehicle() {
	bmw := objects.Bike{}
	fmt.Print("BMW goes: ")
	bmw.Startup()

	lambo := objects.Car{}
	fmt.Print("Lambo goes: ")
	lambo.Startup()
}
func shape() {
	rect := objects.Rectangle{
		Height: 2,
		Width:  4,
	}

	circ := objects.Circle{
		Radius: 2,
	}

	fmt.Println("Rectangle Area: ", rect.Area())
	fmt.Println("Circle Area: ", circ.Area())
	fmt.Println("Rectangle Perimeter: ", rect.Perimeter())
	fmt.Println("Circle Perimeter: ", circ.Perimeter())
}

func calculator() {
	var input int
	fmt.Print("Choices:\n1. Addition\n2. Subtraction\n3. Multiplication\n4. Division\n0. Exit\n")
	fmt.Print("\nEnter Choice: ")
	fmt.Scanln(&input)

	calc := objects.Calculator{}
	if input <= 4 && input != 0 {

		var a, b float64
		fmt.Print("Enter a number: ")
		fmt.Scanln(&a)
		fmt.Print("Enter another number: ")
		fmt.Scanln(&b)
		fmt.Println()

		switch input {
		case 1:
			funcCall, err := calc.Add(a, b)
			if err != nil {
				fmt.Println("Error in function call", err)
			}
			fmt.Printf("Result of Add: %f \n", funcCall)
			fmt.Println(calc.Describe("add"))
			calculator()
		case 2:
			funcCall, err := calc.Subtract(a, b)
			if err != nil {
				fmt.Println("Error in function call", err)
			}
			fmt.Printf("Result of Subtract: %f \n", funcCall)
			fmt.Println(calc.Describe("subtract"))
			calculator()
		case 3:
			funcCall, err := calc.Multiply(a, b)
			if err != nil {
				fmt.Println("Error in function call", err)
			}
			fmt.Printf("Result of Multiply: %f \n", funcCall)
			fmt.Println(calc.Describe("multiply"))
			calculator()
		case 4:
			funcCall, err := calc.Divide(a, b)
			if err != nil {
				fmt.Println("Error in function call", err)
			}
			fmt.Printf("Result of Divide: %f \n", funcCall)
			fmt.Println(calc.Describe("divide"))
			calculator()
		case 0:
			break
		default:
			log.Fatal("Incorrect input: ", input)
		}
		if input > 4 {
			fmt.Println("Incorrect input: ", input)
			calculator()
		}
	}

}
