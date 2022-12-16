package main

import (
	"fmt"
	"project1/helper"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTicket uint = 50

// var bookings [50]string
// var bookings []string
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTicket uint
}

func main() {

	greetUser()

	for {
		firstName, lastName, email, userTicket := getInput()

		isvalidName, isvalidEmail, isvalidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTicket)

		if isvalidEmail && isvalidName && isvalidTicketNumber {

			if remainingTicket == 0 {

				println("we are out of ticket please try again! next year")
				break
			}

			bookTicket(userTicket, firstName, lastName, conferenceName, email)
			sendTicket(userTicket, firstName, lastName, email)

			var first = getFirstName()

			fmt.Printf("The first name of the bookings are: %v\n", first)

		} else {

			if !isvalidEmail {
				fmt.Println("Your email address is empty")
			}

			if !isvalidName {
				fmt.Println("Your Name is empty")
			}

			if !isvalidTicketNumber {
				fmt.Println("Empty ticket is empty!")
			}
		}
	}

}

func greetUser() {

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still remaining ", conferenceTickets, remainingTicket)
	fmt.Println("Get your ticket here to attend")

}

func getFirstName() []string {

	firstNames := []string{}

	for _, booking := range bookings {
		//var name = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)

	}

	return firstNames
}

func getInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your first Name: ")

	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter Number of ticket: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lastName string, conferenceName string, email string) {

	remainingTicket = remainingTicket - userTicket

	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOfTicket: userTicket,
	}

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTicket"] = strconv.FormatUint(uint64(userTicket), 10)

	bookings = append(bookings, userData)

	//now print the list of booking
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v ticket. You will receive a confirmation email at %v \n", firstName, lastName, conferenceName, email)

	fmt.Printf("%v tickets remaining for %v \n", remainingTicket, conferenceName)
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("#####################")
}
