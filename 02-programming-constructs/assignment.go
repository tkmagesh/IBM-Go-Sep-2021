package main

import "fmt"

func main() {
	var userChoice, n1, n2, result int
	for {
		fmt.Println("Menu")
		fmt.Println("================================================================")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")
		fmt.Println("Enter the choice :")
		fmt.Scanf("%d", &userChoice)
		switch userChoice {
		case 1, 2, 3, 4:
			fmt.Println("Enter the two numbers :")
			fmt.Scanf("%d %d", &n1, &n2)
			switch userChoice {
			case 1:
				result = n1 + n2
				fmt.Println("Addition of two numbers is :", result)
			case 2:
				result = n1 - n2
				fmt.Println("Subtraction of two numbers is :", result)
			case 3:
				result = n1 * n2
				fmt.Println("Multiplication of two numbers is :", result)
			case 4:
				result = n1 / n2
				fmt.Println("Division of two numbers is :", result)
			}
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
