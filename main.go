package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println(conferenceTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Println(conferenceName)
}