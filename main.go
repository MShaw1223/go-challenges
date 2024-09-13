package main

import (
	"fmt"
	"log"

	datastructs "github.com/MShaw1223/go-challenges.git/data-structures"
	filesystem "github.com/MShaw1223/go-challenges.git/file-system"
	objects "github.com/MShaw1223/go-challenges.git/objects"
)

func main(){
	var input int
	fmt.Println("Choose an which solutions to run from the following:\n1) Data structures \n2) File system\n3) Objects")
	fmt.Print("\nEnter Input here: ")
	_, err := fmt.Scanln(&input)
	fmt.Println()

	if err != nil {
		log.Fatal("Input Error: ",err)
	}
	
	switch {
		case input == 1:
			fmt.Println("Datastructs module Starting...\n")
			datastructs.Array()
			datastructs.Tuple()
			datastructs.Hash_map()
			fmt.Println("\nDatastructs module complete...\n")
		case input == 2:
			fmt.Println("Filesystem module Starting...\n")
			filesystem.FileSystem()
			fmt.Println("\nFilesystem module complete...\n")
		case input ==3:
			fmt.Println("Bank module Starting...\n")
			bank()
			fmt.Println("\nBank module complete...\n")
			fmt.Println("Vehicle module Starting...\n")
			vehicle()
			fmt.Println("\nVehicle module complete...\n")
			fmt.Println("Shapes module starting...\n")
			shape();
			fmt.Println("\nShapes module complete...\n")
		default: 
			fmt.Println("Incorrect input\n")
			main()
	}	
}

func bank(){
	// make new person
	person:= objects.BankAccount{
		AccountNum: 1,
		Balance: 200.32,
	}

	// use methods in person
	person.See_Balance()

	person.Deposit(20000.68)
	person.See_Balance()

	person.Withdraw(15000)
	person.See_Balance()
}

func vehicle() {
	bmw:=objects.Bike{}
	fmt.Print("BMW goes: ")
	bmw.Startup()

	lambo:=objects.Car{}
	fmt.Print("Lambo goes: ")
	lambo.Startup()
}
func shape(){
	rect:= objects.Rectangle {
		Height: 2,
		Width: 4,
	}

	circ:=objects.Circle {
		Radius: 2,
	}

	fmt.Println("Rectangle Area: ", rect.Area());
	fmt.Println("Circle Area: ", circ.Area());
	fmt.Println("Rectangle Perimeter: ", rect.Perimeter());
	fmt.Println("Circle Perimeter: ", circ.Perimeter());
}

