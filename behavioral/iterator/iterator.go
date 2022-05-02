package main

/*迭代器*/

type iterator interface {
	hasNext() bool
	getNext() *user
}
