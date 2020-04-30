package pack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRowAddSmall(t *testing.T) {

	r := NewRow(3, 2)

	r = r.Add(Entry{
		Size: Small,
	})

	expected := Row{
		Coverage: [][]bool{[]bool{true, false, false}, []bool{false, false, false}},
		C:        1,
		W:        3,
		H:        2,
		Entries: []Entry{
			Entry{Size: Small, X: 0, Y: 0},
		},
	}

	assert.Equal(t, expected, r)

	r = r.Add(Entry{
		Size: Small,
	})

	expected = Row{
		Coverage: [][]bool{[]bool{true, true, false}, []bool{false, false, false}},
		C:        2,
		W:        3,
		H:        2,
		Entries: []Entry{
			Entry{Size: Small, X: 0, Y: 0},
			Entry{Size: Small, X: 1, Y: 0},
		},
	}

	assert.Equal(t, expected, r)
}

func TestRowAddMedium(t *testing.T) {

	r := NewRow(3, 2)

	r = r.Add(Entry{
		Size: Medium,
	})

	expected := Row{
		Coverage: [][]bool{[]bool{true, false, false}, []bool{true, false, false}},
		C:        2,
		W:        3,
		H:        2,
		Entries: []Entry{
			Entry{Size: Medium, X: 0, Y: 0},
		},
	}

	assert.Equal(t, expected, r)

	r = r.Add(Entry{
		Size: Medium,
	})

	expected = Row{
		Coverage: [][]bool{[]bool{true, true, false}, []bool{true, true, false}},
		C:        4,
		W:        3,
		H:        2,
		Entries: []Entry{
			Entry{Size: Medium, X: 0, Y: 0},
			Entry{Size: Medium, X: 1, Y: 0},
		},
	}

	assert.Equal(t, expected, r)
}

func TestRowAddLarge(t *testing.T) {

	r := NewRow(3, 2)

	r = r.Add(Entry{Size: Large})

	expected := Row{
		Coverage: [][]bool{[]bool{true, true, false}, []bool{true, true, false}},
		C:        4,
		W:        3,
		H:        2,
		Entries: []Entry{
			Entry{Size: Large, X: 0, Y: 0},
		},
	}

	assert.Equal(t, expected, r)

	r = r.Add(Entry{Size: Medium})

	expected = Row{
		Coverage: [][]bool{[]bool{true, true, true}, []bool{true, true, true}},
		C:        6,
		W:        3,
		H:        2,
		Entries: []Entry{
			Entry{Size: Large, X: 0, Y: 0},
			Entry{Size: Medium, X: 2, Y: 0},
		},
	}

	assert.Equal(t, expected, r)
}

func TestRowAddXLarge(t *testing.T) {

	r := NewRow(3, 3)

	r = r.Add(Entry{Size: XLarge})

	expected := Row{
		Coverage: [][]bool{[]bool{true, true, false}, []bool{true, true, false}, []bool{true, true, false}},
		C:        6,
		W:        3,
		H:        3,
		Entries: []Entry{
			Entry{Size: XLarge, X: 0, Y: 0},
		},
	}

	assert.Equal(t, expected, r)

	r = r.Add(Entry{Size: Medium})

	expected = Row{
		Coverage: [][]bool{[]bool{true, true, true}, []bool{true, true, true}, []bool{true, true, false}},
		C:        8,
		W:        3,
		H:        3,
		Entries: []Entry{
			Entry{Size: XLarge, X: 0, Y: 0},
			Entry{Size: Medium, X: 2, Y: 0},
		},
	}

	assert.Equal(t, expected, r)
}

func TestRowPack3x3(t *testing.T) {
	r := NewRow(3, 3)

	es := []Entry{
		Entry{Size: Small},
		Entry{Size: XLarge},
		Entry{Size: Medium},
	}

	q := []Entry{}

	r, es = r.Pack(es)

	expected := Row{
		Coverage: [][]bool{[]bool{true, true, true}, []bool{true, true, true}, []bool{true, true, true}},
		C:        9,
		W:        3,
		H:        3,
		Entries: []Entry{
			Entry{Size: Small, X: 0, Y: 0},
			Entry{Size: XLarge, X: 1, Y: 0},
			Entry{Size: Medium, X: 0, Y: 1},
		},
	}

	assert.Equal(t, expected, r)
	assert.Equal(t, q, es)

}

func TestRowPack(t *testing.T) {
	r := NewRow(3, 2)

	es := []Entry{
		Entry{Size: Large},
		Entry{Size: Small},
		Entry{Size: Medium},
		Entry{Size: Small},
		Entry{Size: Large},
		Entry{Size: Small},
		Entry{Size: Medium},
		Entry{Size: Small},
		Entry{Size: Large},
		Entry{Size: Small},
		Entry{Size: Medium},
		Entry{Size: Small},
	}

	q := []Entry{
		Entry{Size: Medium},
		Entry{Size: Large},
		Entry{Size: Small},
		Entry{Size: Medium},
		Entry{Size: Small},
		Entry{Size: Large},
		Entry{Size: Small},
		Entry{Size: Medium},
		Entry{Size: Small},
	}

	r, es = r.Pack(es)

	expected := Row{
		Coverage: [][]bool{[]bool{true, true, true}, []bool{true, true, true}},
		C:        6,
		W:        3,
		H:        2,
		Entries: []Entry{
			Entry{Size: Large, X: 0, Y: 0},
			Entry{Size: Small, X: 2, Y: 0},
			Entry{Size: Small, X: 2, Y: 1},
		},
	}

	assert.Equal(t, expected, r)
	assert.Equal(t, q, es)
}

func BenchmarkPack(b *testing.B) {
	r := NewRow(3, 2)

	es := []Entry{
		Entry{Size: Large},
		Entry{Size: Small},
		Entry{Size: Medium},
		Entry{Size: Small},
	}

	for n := 0; n < b.N; n++ {
		r.Pack(es)
	}
}
