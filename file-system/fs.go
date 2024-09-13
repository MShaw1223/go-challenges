package filesystem

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func FileSystem(){
	// Display options
	fmt.Println("C - Create a file\nD - Delete a file\nA - Append text to a file\nQ - Quit the program")
	var choice string
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	var upper_choice = strings.ToUpper(choice)

	// Call function based on input
	switch {
		case upper_choice == "C":
			CreateFile();
			FileSystem();
		case upper_choice == "Q":
			break;
		default:
			fmt.Println("Incorrect input");
			FileSystem();
	}
}

func CreateFile(){
	file, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

func DeleteFile(){
	err := os.Remove("example.txt");
	if err != nil {
		log.Fatal(err)
	}
}
