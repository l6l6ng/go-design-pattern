package main

/*备忘录*/

type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}
