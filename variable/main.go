package main

import "fmt"

const LoginToken string = "asdnasdsad" // First letter uppercase means its public

func main() {
	var username string = "Prabhat"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.455654234234234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int //default is 0
	fmt.Println(anotherVariable)

	// Implicit type
	var website = "Prabhat.singh.com"

	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//website = 3 //not allowed as it is already treated as string implicitly

	//NO var style :

	numberOfUser := 30000 // This style is allowed only inside method
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)

	fmt.Printf("LoginToken %T", LoginToken)

}
