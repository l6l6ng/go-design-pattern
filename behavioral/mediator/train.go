package main

/*组件*/

type train interface {
	arrive()
	depart()
	permitArrival()
}
