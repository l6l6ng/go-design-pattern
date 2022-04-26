package main

import "fmt"

/*具体实施*/

type epson struct {}

func (p *epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}
