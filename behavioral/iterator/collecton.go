package main

/*集合*/

type collection interface {
	createIterator() iterator
}
