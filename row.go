package pack

// Row is a 3x2 grid fully (or almost) covered by entries
type Row struct {
	Coverage [2][3]bool
	Entries  []Entry

	Covered int
}

// Add adds (:P) (or tries to) an Entry to the Row. If the Entry cannot fit the Row is left untouched. It's chainable.
func (r Row) Add(e Entry) Row {
	if e.Size > 6-r.Covered {
		return r
	}

	switch e.Size {
	case SMALL:
		for i, rr := range r.Coverage {
			for j, v := range rr {
				if !v {
					r.Entries = append(r.Entries, e.Pos(j, i))
					r.Covered++
					r.Coverage[i][j] = true
					return r
				}
			}
		}
	case MEDIUM:
		for i, _ := range r.Coverage[0] {
			if !r.Coverage[0][i] && !r.Coverage[1][i] {
				r.Entries = append(r.Entries, e.Pos(i, 0))
				r.Covered += 2
				r.Coverage[0][i] = true
				r.Coverage[1][i] = true
				return r
			}
		}
	case LARGE:
		if !r.Coverage[0][0] && !r.Coverage[0][1] && !r.Coverage[1][0] && !r.Coverage[1][1] {
			r.Entries = append(r.Entries, e.Pos(0, 0))
			r.Covered += 4
			r.Coverage[0][0] = true
			r.Coverage[1][0] = true
			r.Coverage[1][1] = true
			r.Coverage[0][1] = true
			return r
		}
		if !r.Coverage[0][1] && !r.Coverage[0][2] && !r.Coverage[1][1] && !r.Coverage[1][2] {
			r.Entries = append(r.Entries, e.Pos(1, 0))
			r.Covered += 4
			r.Coverage[0][1] = true
			r.Coverage[1][1] = true
			r.Coverage[1][2] = true
			r.Coverage[0][2] = true
			return r
		}
	}

	return r
}

// Add Pack iterates over all the provided an Entries and tries to add them to the Row. If the Entry cannot fit in the Row it's skipped and later returned. It's chainable.
func (r Row) Pack(e []Entry) (Row, []Entry) {
	q := []Entry{}
	p := r.Covered
	for _, v := range e {
		r = r.Add(v)
		if p == r.Covered {
			q = append(q, v)
		} else {
			p = r.Covered
		}
	}
	return r, q
}
