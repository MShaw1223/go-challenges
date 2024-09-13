package main

import (
	"fmt"

	filesystem "github.com/MShaw1223/go-challenges.git/file-system"
	objects "github.com/MShaw1223/go-challenges.git/objects"
)

func main(){
	fmt.Println("Bank module Starting...")
	// bank()
	fmt.Println("Bank module complete...")
	fmt.Println("Vehicle module Starting...")
	// vehicle()
	fmt.Println("Vehicle module complete...")
	fmt.Println("Datastructs module Starting...")
	// datastructs.Array()
	// datastructs.Tuple()
	// datastructs.Hash_map()
	fmt.Println("Datastructs module complete...")
	fmt.Println("Filesystem module Starting...")
	filesystem.FileSystem()
	fmt.Println("Filesystem module complete...")
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
	bmw.Startup()

	lambo:=objects.Car{}
	lambo.Startup()

}
