package main

import "fmt"

func main() {
	hpPrinter := &Hp{}
	epson := &Epson{}

	macComputer := &Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epson)
	macComputer.Print()
	fmt.Println()

	winComputer := &Windows{}
	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epson)
	winComputer.Print()
	fmt.Println()
}
