package set

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(vals ...T) {
	for _, v := range vals {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}
