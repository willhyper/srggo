package srg

func (s *Srg) Copy() *Srg{

	ss := NewSrg(s.v, s.k, s.l, s.u)
	ss.filled = s.filled
	ss.toFill = s.toFill
	ss.Matrix = append([]bool(nil), s.Matrix...)

	ss.SymSquare = make([][]bool, len(s.SymSquare))
	for i, v:= range s.SymSquare{
		ss.SymSquare[i] = append([]bool(nil), v...)
	}

	return ss
}