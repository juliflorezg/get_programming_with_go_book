package main

type turtle struct {
	x, y int
}

func up(t *turtle) {
	t.y--
}
func down(t *turtle) {
	t.y++
}
func left(t *turtle) {
	t.x--
}
func right(t *turtle) {
	t.x++
}
