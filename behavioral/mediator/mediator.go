package main

/*中介者接口*/

type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}
