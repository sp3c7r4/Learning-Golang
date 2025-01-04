package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	var firstName string;
	var lastName string;
	var email string;
	var userTickets int;
	var remainingTickets uint = 50;
	//Array type
	var bookings []string;
	
	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	for remainingTickets > 0 && len(bookings) < 50 {
		//Asking for user Input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Enter the number of tickets you wanna buy: ")
		fmt.Scan(&userTickets)
	
		if userTickets < int(remainingTickets)  {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, firstName + " " + lastName)
	
			fmt.Printf("Thank you %v for booking %v tickets. You will recieve a confirmation email at %v \n", bookings[0], userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
			var firstNames = []string{}
			for _ , booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These first names  of bookings are: %v \n", firstNames)
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year")
				break;
			}
		} else if userTickets == int(remainingTickets) {
			fmt.Printf("Same Tickets")
		} else {
			fmt.Printf("We only have %v tickets ramaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}
}