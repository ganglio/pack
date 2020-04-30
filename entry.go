package pack

type Size struct {
	W int
	H int
}

func (s Size) Size() int {
	return s.W * s.H
}

var (
	Small  = Size{1, 1}
	Medium = Size{1, 2}
	Large  = Size{2, 2}
	XLarge = Size{2, 3}
)

// Entry represent an entry to be packed (or already packed) in a Row
type Entry struct {
	X    int
	Y    int
	Size Size
	Data interface{}
}

func (e Entry) Pos(x, y int) Entry {
	e.X = x
	e.Y = y
	return e
}
