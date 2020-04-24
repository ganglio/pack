package pack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRowAddSmall(t *testing.T) {

	r := Row{}

	r = r.Add(Entry{
		Size: SMALL,
	})

	expected := Row{
		Coverage: [2][3]bool{[3]bool{true, false, false}, [3]bool{false, false, false}},
		Covered:  1,
		Entries: []Entry{
			Entry{Size: SMALL, X: 0, Y: 0},
		},
	}

	assert.Equal(t, expected, r)

	r = r.Add(Entry{
		Size: SMALL,
	})

	expected = Row{
		Coverage: [2][3]bool{[3]bool{true, true, false}, [3]bool{false, false, false}},
		Covered:  2,
		Entries: []Entry{
			Entry{Size: SMALL, X: 0, Y: 0},
			Entry{Size: SMALL, X: 1, Y: 0},
		},
	}

	assert.Equal(t, expected, r)
}

func TestRowAddMedium(t *testing.T) {

	r := Row{}

	r = r.Add(Entry{
		Size: MEDIUM,
	})

	expected := Row{
		Coverage: [2][3]bool{[3]bool{true, false, false}, [3]bool{true, false, false}},
		Covered:  2,
		Entries: []Entry{
			Entry{Size: MEDIUM, X: 0, Y: 0},
		},
	}

	assert.Equal(t, expected, r)

	r = r.Add(Entry{
		Size: MEDIUM,
	})

	expected = Row{
		Coverage: [2][3]bool{[3]bool{true, true, false}, [3]bool{true, true, false}},
		Covered:  4,
		Entries: []Entry{
			Entry{Size: MEDIUM, X: 0, Y: 0},
			Entry{Size: MEDIUM, X: 1, Y: 0},
		},
	}

	assert.Equal(t, expected, r)
}

func TestRowAddLarge(t *testing.T) {

	r := Row{}

	r = r.Add(Entry{Size: LARGE})

	expected := Row{
		Coverage: [2][3]bool{[3]bool{true, true, false}, [3]bool{true, true, false}},
		Covered:  4,
		Entries: []Entry{
			Entry{Size: LARGE, X: 0, Y: 0},
		},
	}

	assert.Equal(t, expected, r)

	r = r.Add(Entry{Size: MEDIUM})

	expected = Row{
		Coverage: [2][3]bool{[3]bool{true, true, true}, [3]bool{true, true, true}},
		Covered:  6,
		Entries: []Entry{
			Entry{Size: LARGE, X: 0, Y: 0},
			Entry{Size: MEDIUM, X: 2, Y: 0},
		},
	}

	assert.Equal(t, expected, r)
}

func TestRowPack(t *testing.T) {
	r := Row{}

	es := []Entry{
		Entry{Size: LARGE},
		Entry{Size: SMALL},
		Entry{Size: MEDIUM},
		Entry{Size: SMALL},
		Entry{Size: LARGE},
		Entry{Size: SMALL},
		Entry{Size: MEDIUM},
		Entry{Size: SMALL},
		Entry{Size: LARGE},
		Entry{Size: SMALL},
		Entry{Size: MEDIUM},
		Entry{Size: SMALL},
	}

	q := []Entry{
		Entry{Size: MEDIUM},
		Entry{Size: LARGE},
		Entry{Size: SMALL},
		Entry{Size: MEDIUM},
		Entry{Size: SMALL},
		Entry{Size: LARGE},
		Entry{Size: SMALL},
		Entry{Size: MEDIUM},
		Entry{Size: SMALL},
	}

	r, es = r.Pack(es)

	expected := Row{
		Coverage: [2][3]bool{[3]bool{true, true, true}, [3]bool{true, true, true}},
		Covered:  6,
		Entries: []Entry{
			Entry{Size: LARGE, X: 0, Y: 0},
			Entry{Size: SMALL, X: 2, Y: 0},
			Entry{Size: SMALL, X: 2, Y: 1},
		},
	}

	assert.Equal(t, expected, r)
	assert.Equal(t, q, es)
}

func BenchmarkPack(b *testing.B) {
	r := Row{}

	es := []Entry{
		Entry{Size: LARGE},
		Entry{Size: SMALL},
		Entry{Size: MEDIUM},
		Entry{Size: SMALL},
	}

	for n := 0; n < b.N; n++ {
		r.Pack(es)
	}
}
