package main

/*具体元件*/

type rectangle struct {
	l int
	b int
}

func (t *rectangle) accept(v visitor) {
	v.visitForRectangle(t)
}

func (t *rectangle) getType() string {
	return "Rectangle"
}
