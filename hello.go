package main

import (
	"fmt"
	"golang/helper"
	"sync"
	"time"
)

var conferenceName = "Seminar Cybersec"

const conferenceTicket uint = 50

var remainingTicket uint = conferenceTicket

var bookings = make([]UserData, 0)

type UserData struct {
	firstName  string
	lastName   string
	email      string
	userTicket uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	firstName, lastName, email, userTicket := getUserInput()

	isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTicket)

	if isValidEmail && isValidName && isValidTicket {

		bookTicket(userTicket, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTicket, firstName, lastName, email)

		firstNames := getFirstName()
		fmt.Printf("Berikut Ticket Booking anda: %v\n", firstNames)

		noTicketRemaining := remainingTicket == 0
		if noTicketRemaining {
			fmt.Printf("%v Ticket Seminar tersebut soldout!\n", conferenceName)
		}

	} else {
		if !isValidName {
			fmt.Println("Nama terlalu pendek")
		}

		if !isValidEmail {
			fmt.Println("Email yang anda masukkan invalid")
		}

		if !isValidTicket {
			fmt.Println("Anda memasukkan no ticket yang salah")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get Your Ticket\n")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Masukkan nama depan anda:")
	fmt.Scan(&firstName)

	fmt.Println("Masukkan nama belakang anda:")
	fmt.Scan(&lastName)

	fmt.Println("Masukkan email address anda")
	fmt.Scan(&email)

	fmt.Println("Masukkan ticket yang anda pilih")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTicket -= userTicket

	var userData = UserData{
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		userTicket: userTicket,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank You %v %v for Booking %v Tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v Tickets Remaining at %v Booking Application\n", remainingTicket, conferenceName)
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var waktu = time.Now().Format("02/01/2006 15:04:05")
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("########################################################################################")
	fmt.Printf("Sending ticket:\n    %v to email address %v at %v\n", ticket, email, waktu)
	fmt.Println("########################################################################################")
	wg.Done()
}