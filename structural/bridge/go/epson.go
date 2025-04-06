package main

import "fmt"

type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by an Epson Printer")
}
