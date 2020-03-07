package main

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

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	set := NewSet()
	for {
		set.Add(n)
		n = sumOfSquaredDigits(n)
		if set.Contains(n) {
			return false
		}
		if n == 1 {
			return true
		}
	}
}

func sumOfSquaredDigits(transitiveNumber int) int {
	sum, num := 0, 0
	for transitiveNumber != 0 {
		num = transitiveNumber % 10
		sum += num * num
		transitiveNumber /= 10
	}
	return sum
}
