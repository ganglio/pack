package pack

type coverage [][]bool

type Row struct {
	Coverage coverage
	Entries  []Entry

	W int
	H int
	C int
}

func NewRow(w, h int) Row {

	r := make([][]bool, h)
	for k, _ := range r {
		r[k] = make([]bool, w)
	}

	return Row{
		Coverage: r,
		W:        w,
		H:        h,
		C:        0,
	}
}

func (s coverage) empty(w, h, i, j int) bool {
	f := false
	for n := 0; n < h; n++ {
		for m := 0; m < w; m++ {
			f = f || s[n+j][m+i]
		}
	}
	return !f
}

func (s coverage) fill(w, h, i, j int) {
	for n := 0; n < h; n++ {
		for m := 0; m < w; m++ {
			s[n+j][m+i] = true
		}
	}
}

// Add adds (:P) (or tries to) an Entry to the Row. If the Entry cannot fit the Row is left untouched. It's chainable.
func (r Row) Add(e Entry) Row {
	if e.Size.Size() > r.W*r.H-r.C {
		return r
	}

	// Interate over the part of the grid where the entry could be added
	for j := 0; j < r.H-e.Size.H+1; j++ {
		for i := 0; i < r.W-e.Size.W+1; i++ {
			// Iterate over the Entry
			if r.Coverage.empty(e.Size.W, e.Size.H, i, j) {
				r.Entries = append(r.Entries, e.Pos(i, j))
				r.C += e.Size.Size()
				r.Coverage.fill(e.Size.W, e.Size.H, i, j)
				return r
			}
		}
	}

	return r
}

// Add Pack iterates over all the provided an Entries and tries to add them to the Row. If the Entry cannot fit in the Row it's skipped and later returned. It's chainable.
func (r Row) Pack(e []Entry) (Row, []Entry) {
	q := []Entry{}
	p := r.C
	for _, v := range e {
		r = r.Add(v)
		if p == r.C {
			q = append(q, v)
		} else {
			p = r.C
		}
	}
	return r, q
}
