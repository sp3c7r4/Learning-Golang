package main

import "fmt"

func main() {

	var firstName string;
	var lastName string;
	var email string;
	var userTickets int;
	
	//Asking for user Input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets you wanna buy: ")
	fmt.Scan(&userTickets)
	// Formating Variables
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v \n", firstName, lastName, userTickets, email)

}