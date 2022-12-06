package set

var exists = struct{}{}

type Set struct {
	m map[int]struct{}
}

func set() *Set {
	s := &Set{}
	s.m = make(map[int]struct{})
	return s
}

func (s *Set) add(value int) {
	s.m[value] = exists
}

func (s *Set) delete(value int) {
	delete(s.m, value)
}

func (s *Set) contains(value int) bool {
	_, ok := s.m[value]
	return ok
}
