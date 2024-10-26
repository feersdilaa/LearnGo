package main

import (
	"fmt"
)

func scanInput() (string, string) {
	var nama, nrp string
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan NRP: ")
	fmt.Scanln(&nrp)
	return nama, nrp
}

func printMessage(nama, nrp string) {
	fmt.Println("Nama:", nama)
	fmt.Println("NRP:", nrp)
	fmt.Println(nama, "Dengan NRP", nrp, "Telah Login")
}

func doScanAndPrint() {

	nama, nrp := scanInput()
	
	printMessage(nama, nrp)
}

func main() {
	fmt.Println("Program Sederhana")

	doScanAndPrint()
}
