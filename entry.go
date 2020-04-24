package pack

const (
	SMALL  = 1
	MEDIUM = 2
	LARGE  = 4
)

// Entry represent an entry to be packed (or already packed) in a Row
type Entry struct {
	X    int
	Y    int
	Size int
	Data interface{}
}

func (e Entry) Pos(x, y int) Entry {
	e.X = x
	e.Y = y
	return e
}
