package main

import "fmt"

func main() {
	fmt.Println()

	client := &Client{}
	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	fmt.Printf("\n******************\n\n")

	windowsMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{windowsMachine}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

	fmt.Println()
}
