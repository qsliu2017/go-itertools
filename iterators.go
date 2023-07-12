package itertools

type Iterator[T any] interface {
	Next() (T, bool)
}

type infIterator int

func (i *infIterator) Next() (int, bool) {
	*i++
	return int(*i), true
}

// Inf returns an infinite iterator of int.
func Inf() Iterator[int] { return new(infIterator) }

type repeatIterator[T any] struct{ t T }

func (i *repeatIterator[T]) Next() (T, bool) { return i.t, true }

func Repeat[T any](t T) Iterator[T] { return &repeatIterator[T]{t} }

type sliceIterator[T any] struct {
	slice []T
	index int
}

func (i *sliceIterator[T]) Next() (T, bool) {
	var t T
	if i.index < len(i.slice) {
		t = i.slice[i.index]
		i.index++
		return t, true
	}
	return t, false
}

func OfSlice[T any](slice []T) Iterator[T] { return &sliceIterator[T]{slice, 0} }

type chanIterator[T any] struct {
	ch <-chan T
}

func (i *chanIterator[T]) Next() (T, bool) {
	t, ok := <-i.ch
	return t, ok
}

func OfChan[T any](ch <-chan T) Iterator[T] { return &chanIterator[T]{ch} }

type mapIterator[T, U any] struct {
	iter Iterator[T]
	f    func(T) U
}

func (i *mapIterator[T, U]) Next() (U, bool) {
	if t, ok := i.iter.Next(); ok {
		return i.f(t), true
	}
	var u U
	return u, false
}

func Map[T, U any](iter Iterator[T], f func(T) U) Iterator[U] {
	return &mapIterator[T, U]{iter, f}
}

type flatMapIterator[T, U any] struct {
	tIter Iterator[T]
	f     func(T) Iterator[U]
	uIter Iterator[U]
}

func (i *flatMapIterator[T, U]) Next() (U, bool) {
	for {
		if i.uIter == nil {
			if t, ok := i.tIter.Next(); ok {
				i.uIter = i.f(t)
			} else {
				var u U
				return u, false
			}
		}
		if u, ok := i.uIter.Next(); ok {
			return u, true
		} else {
			i.uIter = nil
		}
	}
}

func FlatMap[T, U any](iter Iterator[T], f func(T) Iterator[U]) Iterator[U] {
	return &flatMapIterator[T, U]{iter, f, nil}
}

type filterIterator[T any] struct {
	iter Iterator[T]
	take func(T) bool
}

func (i *filterIterator[T]) Next() (T, bool) {
	for {
		t, ok := i.iter.Next()
		if !ok {
			return t, false
		}
		if i.take(t) {
			return t, true
		}
	}
}

func Filter[T any](iter Iterator[T], take func(T) bool) Iterator[T] {
	return &filterIterator[T]{iter, take}
}

type takeIterator[T any] struct {
	iter Iterator[T]
	n    int
	i    int
}

func (i *takeIterator[T]) Next() (T, bool) {
	if i.i < i.n {
		i.i++
		return i.iter.Next()
	}
	var t T
	return t, false
}

func Take[T any](iter Iterator[T], n int) Iterator[T] {
	return &takeIterator[T]{iter, n, 0}
}
