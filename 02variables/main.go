package main

import "fmt"

//putting the first letter as captial intializes that variable as a public variable
const LoginToken string = "adityam"

func main() {
	var username string = "Aditya"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)
	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)
	var smallFloat float64 = 256.32424234234234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)
	//defualt values and aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T\n", anotherVar)
	//implicit way of declaring a variable
	var website = "Thesocialknowledge"
	fmt.Println(website)
	//no var style
	numberOfUsers := 23444
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T\n", LoginToken)
}
