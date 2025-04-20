package main

func main() {
	tv := &Tv{}
	onCommand := &OnCommand{tv}
	offCommand := &OffCommand{tv}

	onButton := &Button{command: onCommand}
	onButton.press()

	offButton := &Button{command: offCommand}
	offButton.press()
}
