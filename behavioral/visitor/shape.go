package main

/*元件*/

type shape interface {
	getType() string
	accept(visitor)
}
