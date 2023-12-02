package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welsome := "Welcome to user input"
	fmt.Println(welsome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")
	// comma ok syntax or error okay syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of the rating is %T", input)
}
