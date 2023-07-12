package itertools

func ForEach[T any](iter Iterator[T], f func(T)) {
	for t, ok := iter.Next(); ok; t, ok = iter.Next() {
		f(t)
	}
}

func ForEachIndexed[T any](iter Iterator[T], f func(int, T)) {
	i := 0
	ForEach(iter, func(t T) {
		f(i, t)
		i++
	})
}

func ToChan[T any](iter Iterator[T]) <-chan T {
	ch := make(chan T)
	go func() {
		ForEach(iter, func(t T) { ch <- t })
		close(ch)
	}()
	return ch
}

func Reduce[T, R any](iter Iterator[T], inital R, reducer func(R, T) R) R {
	r := inital
	ForEach(iter, func(t T) {
		r = reducer(r, t)
	})
	return r
}

func ToSlice[T any](iter Iterator[T]) []T {
	return Reduce(
		iter,
		make([]T, 0),
		func(slice []T, t T) []T { return append(slice, t) },
	)
}

// Set represents a set of T.
type Set[T comparable] map[T]struct{}

func GroupBy[T any, K, V comparable](iter Iterator[T], keyMapper func(T) K, valueMapper func(T) V) map[K]Set[V] {
	return Reduce(
		iter,
		make(map[K]Set[V]),
		func(m map[K]Set[V], t T) map[K]Set[V] {
			k := keyMapper(t)
			v := valueMapper(t)
			if _, ok := m[k]; !ok {
				m[k] = make(Set[V])
			}
			m[k][v] = struct{}{}
			return m
		},
	)
}
