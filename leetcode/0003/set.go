package main

type set[T comparable] map[T]struct{}

func (s set[T]) Add(key T) {
	s[key] = struct{}{}
}

func (s set[T]) Has(key T) bool {
	_, exists := s[key]
	return exists
}

func (s set[T]) Remove(key T) {
	delete(s, key)
}
