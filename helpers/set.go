package helpers

type Set[T comparable] struct {
	values map[T]struct{}
}

func NewSet[T comparable](values []T) Set[T] {
	set := Set[T]{values: make(map[T]struct{})}
	for _, value := range values {
		set.values[value] = struct{}{}
	}
	return set
}

func (s Set[T]) Exists(item T) bool {
	_, ok := s.values[item]
	return ok
}
