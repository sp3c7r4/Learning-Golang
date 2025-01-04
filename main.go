package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println(conferenceTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println(conferenceName)

	var userName string;
	var userTickets int;
	
	userName = "Tom"
	userTickets = 9
	fmt.Printf("User %v booked %v tickers. \n", userName, userTickets)
}