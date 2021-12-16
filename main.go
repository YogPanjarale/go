//import all the packages
package main

import (
	"fmt"
)

func main() {
	//Take input from the user to run which example
	fmt.Println("Enter 1 for Hello World")
	fmt.Println("Enter 2 for Multiplication Table")
	fmt.Println("Enter 3 for Number Guessing Game")
	var input int
	fmt.Scanf("%d", &input)

	switch input {
	case 1:
		fmt.Println("Press Enter to continue")
		hello_world()
	case 2:
		fmt.Println("Press Enter to continue")
		multiplication_table()
	case 3:
		fmt.Println("Press Enter to continue")
		number_guessing_game()
	default:
		fmt.Println("Invalid Input")
	}
}

