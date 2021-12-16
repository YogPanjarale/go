//make a cli game which take a number and print the multiplication table of that number
package main
import (
	"fmt"
	"os"
)
func multiplication_table() {
	fmt.Scanln()
	fmt.Print("Enter a number: ")
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", input, i, input*i)
	}
}