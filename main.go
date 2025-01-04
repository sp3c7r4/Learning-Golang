package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	var firstName string;
	var lastName string;
	var email string;
	var userTickets int;
	var remainingTickets uint = 50;
	//Array type
	var bookings []string;
	
	//Asking for user Input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets you wanna buy: ")
	fmt.Scan(&userTickets)
	
	remainingTickets = remainingTickets - uint(userTickets)
	bookings = append(bookings, firstName + " " + lastName)
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array Type: %T\n", bookings)
	fmt.Printf("Array Size: %T\n", len(bookings))
	
	// Formating Variables
	fmt.Printf("Thank you %v for booking %v tickets. You will recieve a confirmation email at %v \n", bookings[0], userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}