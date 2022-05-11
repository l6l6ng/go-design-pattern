package main

import "fmt"

/*具体状态*/

type noItemState struct {
	vendingMachine *vendingMachine
}

func (i *noItemState) requestItem() error {
	return fmt.Errorf("Item out of  stock")
}

func (i *noItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *noItemState) insertMoney(money int) error{
	return fmt.Errorf("Item out of  stock")
}

func (i *noItemState) dispenseItem() error {
	return fmt.Errorf("Item out of  stock")
}