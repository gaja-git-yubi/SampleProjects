package main

import "fmt"

func add(a int,b int)(int){
	return a+b
}

func subtract(a int,b int)(int){
	return a-b
}
func multiply(a int,b int )(int){
	return a*b
}

func divide(a,b int)(int){
	return a/b
}
func modulus(a int,b int)(int){
	return a%b
}

func main(){
	var n1, n2 int
	var operator string

	fmt.Print("Enter first number:")
	fmt.Scanln(&n1)

	fmt.Print("Enter Second number:")
	fmt.Scanln(&n2)

	fmt.Print("Enter operation symbol:")
	fmt.Scanln(&operator)

	var result int
	switch operator{

	case "+" :
		result = add(n1,n2)
	case "-" :
		result = subtract(n1,n2)
	case "*" :
		result = multiply(n1,n2)
	case "/" :
		result = divide(n1,n2)
	case "%" :
		result = modulus(n1,n2)
	default:
		fmt.Println("Invalid operator")
		return	
	}
	fmt.Println("Result: ", result)
}