package main

import (
	"booking-app/helper"
	"fmt"
	// "strconv"
)

// variable "Public"
const conferenceTickets int = 50

// Equivalant a
// var conferenceName string = "Go conference"
var conferenceName string = "Go conference" // c'est pas ouff, ça fonctionne pas avec le const et le typage on peut pas changer
// int all numbers positif and negative
// uint only posifiv number
var remainingTickets uint = 50

// [50] 50 elements max
// [] no max
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUser()
	// var bookings = []string{}
	// bookings := []string{}

	// pointeur = &
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)

	for {
		// var userName string
		// var lastName string
		// var email string
		// var userTickets int
		// // ask user for their name
		// fmt.Println("Enter your first name:")
		// fmt.Scan(&userName)
		// fmt.Println("Enter your lastName:")
		// fmt.Scan(&lastName)
		// fmt.Println("Enter your email:")
		// fmt.Scan(&email)
		// fmt.Println("Enter number of tickets:")
		// fmt.Scan(&userTickets)

		userName, lastName, email, userTickets := getUserInput()
		// isValidName := len(userName) >= 2 && len(lastName) >= 2
		// isValidEmail := strings.Contains(email, "@")
		// isValidTickerNumber := userTickets > 0 && userTickets < int(remainingTickets)
		isValidName, isValidEmail, isValidTickerNumber := helper.ValidateUserInpt(userName, lastName, email, userTickets, remainingTickets)
		// isValidCity := city == "singapore" || city == "London"
		// isInvalidCity := city != "singapore" || city != "London"

		if isValidName && isValidEmail && isValidTickerNumber {
			bookTicket(remainingTickets, uint(userTickets), userName, lastName, email, conferenceName)
			sendTicket(userTickets, userName, lastName, email)
			// remainingTickets = remainingTickets - uint(userTickets)
			// bookings = append(bookings, userName+" "+lastName)

			// fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, lastName, userTickets, email)
			// fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("FirstName or LastName you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Emai address is invalid")
			}
			if !isValidTickerNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Println("Your input data in invalid")
			continue // passe au prochain tour
		}
	}

	// city := "London"

	// switch city {
	// case "New York":
	// 	// code
	// case "singapore":
	// 	// execute singapore code
	// case "London","Berlin":
	// 	// execute code
	// case "Hong Kong":
	// 	// code here
	// default:
	// 	// default case
	// 	fmt.Print("No valid city selected")
	// }
}

func greetUser() {
	fmt.Printf("Welcom to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		// firstNames = append(firstNames, booking["firstName"]) avec map
		firstNames = append(firstNames, booking.firstName) // avec strcs
	}
	// fmt.Printf("The first names of bookings are: %v \n", firstNames)
	return firstNames
}

// func validateUserInpt(firstName string, lastName string, email string, userTickets int, remaningTickets uint) (bool, bool, bool) {
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidTickerNumber := userTickets > 0 && userTickets < int(remaningTickets)

// 	return isValidName, isValidEmail, isValidTickerNumber
// }

func getUserInput() (string, string, string, int) {
	var userName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&userName)
	fmt.Println("Enter your lastName:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return userName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, userName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - uint(userTickets)
	// create a map for user

	var userData = UserData{
		firstName:       userName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"] = userName
	// userData["lastName"] = lastName
	// userData["email"] = email

	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Print("###############")
	fmt.Printf("Sending ticket %v to email address %v", ticket, email)
	fmt.Print("###############")

}
