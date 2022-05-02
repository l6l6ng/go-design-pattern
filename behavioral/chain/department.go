package main

/*处理者接口*/

type department interface {
	execute(*patient)
	setNext(department)
}
