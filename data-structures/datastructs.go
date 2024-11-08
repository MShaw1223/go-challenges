package datastructs

import "fmt"

func Array() {
	var input string
	var a [5]string

	for i := 0; i < len(a); i++ {
		fmt.Printf("Enter word number %d of %d here: ", i+1, len(a))
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("Error: %s at index: %d (word num %d)\n", err, i, i+1)
		}
		a[i] = input
	}
	fmt.Println("Array of your words:", a)
}

func Tuple() {
	type tuple struct {
		First  int
		Second string
	}

	var input string
	var int_input int

	fmt.Print("Enter a string: ")
	_, f_err := fmt.Scanln(&input)
	if f_err != nil {
		fmt.Printf("Error in string scanln: %T", f_err)
	}

	fmt.Print("Enter an integer: ")
	_, s_err := fmt.Scanln(&int_input)
	if s_err != nil {
		fmt.Printf("Error in int scanln: %T", s_err)
	}

	t := tuple{First: int_input, Second: input}

	fmt.Println("Tuple:", t)
}

func Hash_map() {
	age_map := make(map[string]int)

	var age int
	var name string

	fmt.Print("Enter an age: ")
	_, a_err := fmt.Scanln(&age)
	if a_err != nil {
		fmt.Println("Error reading age input:", a_err)
	}

	fmt.Print("Enter a name: ")
	_, n_err := fmt.Scanln(&name)
	if n_err != nil {
		fmt.Println("Error reading name input: ", n_err)
	}
	age_map[name] = age

	fmt.Print("Age Map: ", age_map, "\n")
	fmt.Printf("%s's age: %d", name, age_map[name])
}
