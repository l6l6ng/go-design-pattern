package main

import "fmt"

/*具体实施*/

type hp struct{}

func (p *hp) printFile() {
	fmt.Println("Printing by a HP Printer")
}
