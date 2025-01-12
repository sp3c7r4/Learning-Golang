package main;

import (
	"learning-go/helper"
	"fmt"
	"strings"
)

var conferenceName = "Go conference"
var firstName string
var lastName string
var email string
var userTickets int
var remainingTickets = 50

//Array type
func main() {
	var bookings []string
	
	greetUser("Spectra", conferenceName)
	
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

		
		//Validating user inputs
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
			
			fmt.Printf("Thank you %v for booking %v tickets. You will recieve a confirmation email at %v \n", bookings[0], userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			
			//Call function print firstnames
			firstNames := getFirstNames(bookings)
			fmt.Printf("These first names  of bookings are: %v \n", firstNames)
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
			} else {
				if !isValidName {
					fmt.Println("Firstname or lastname you entered is too short ")
				}
				if !isValidEmail {
					fmt.Println("Email address in invalid")
				}
				if !isValidTicketNumber {
					fmt.Println("Number of tickets you entered is invalid")
				}
			}
		}
}

func greetUser(name string, conferenceName string ) {
	fmt.Printf("Welcome %v!\n", name)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames;
}