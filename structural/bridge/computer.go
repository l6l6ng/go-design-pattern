package main

/*抽象*/

type computer interface {
	print()
	setPrinter(printer)
}
