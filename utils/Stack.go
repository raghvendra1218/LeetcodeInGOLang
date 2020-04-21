package utils

var exists = struct{}{}

type set struct {
	m map[int]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[int]struct{})
	return s
}

func (s *set) Add(value int) {
	s.m[value] = exists
}

func (s *set) Remove(value int) {
	delete(s.m, value)
}

func (s *set) Contains(value int) bool {
	_, c := s.m[value]
	return c
}
