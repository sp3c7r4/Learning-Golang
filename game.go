package main;

import (
	"math/rand"
	"fmt"
	"strings"
)

func Game(){
	var userChoice string;
	var attempts uint8;
	var userScore uint8;
	var computerScore uint8;
	var choices = []string{"rock", "paper", "scissors"};
	
	for attempts != 3 {
		randomNumber := rand.Intn(3);
		computerChoice := choices[randomNumber]
		fmt.Println("Enter your choice:")
		fmt.Scanln(&userChoice)
		userChoiceToLower := strings.ToLower(userChoice)
		fmt.Println(userChoiceToLower)

		switch userChoiceToLower {
		case "r", "rock":
			if computerChoice[:1] == userChoiceToLower[:1] {
				fmt.Println("It's a Tie")
			} else if computerChoice == "paper" {
				fmt.Println("You lose")
				computerScore++
			} else if computerChoice == "scissors" {
				fmt.Println("You win!")
				userScore++
			}
		case "p", "paper":
			if computerChoice[:1] == userChoiceToLower[:1] {
				fmt.Println("It's a Tie")
			} else if computerChoice == "rock" {
				fmt.Println("You win!")
				userScore++
			} else if computerChoice == "scissors" {
				fmt.Println("You Lose!")
				computerScore++
			}
		case "s", "scissors":
			if computerChoice[:1] == userChoiceToLower[:1] {
				fmt.Println("It's a Tie")
			} else if computerChoice == "rock" {
				fmt.Println("You Lose!")
				computerScore++
			} else if computerChoice == "paper" {
				fmt.Println("You Win!")
				userScore++
			}
		default:
			fmt.Println("Invalid Input!")
		}
		attempts += 1
	}

	if computerScore > userScore {
		fmt.Println("Computer Wins!")
	} else if computerScore < userScore {
		fmt.Println("User Wins!")
	}
	
	fmt.Printf("🤖: %v 🧑:%v", computerScore, userScore)
}