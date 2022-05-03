package main

type caretaker struct {
	mementoToArray []*memento
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoToArray = append(c.mementoToArray, m)
}

func (c *caretaker) getMemento(index int) *memento {
	return c.mementoToArray[index]
}
