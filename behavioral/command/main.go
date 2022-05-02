package main

/*客户端代码*/

func main() {
	tv := &tv{}

	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &button{
		command: onCommand,
	}

	onButton.press()

	offButton := &button{
		command: offCommand,
	}

	offButton.press()
}
