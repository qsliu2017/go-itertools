package itertools_test

import (
	"testing"

	"github.com/qsliu2017/go-itertools"
	"github.com/stretchr/testify/assert"
)

func TestSliceIterator(t *testing.T) {
	iter := itertools.OfSlice([]int{1, 2, 3, 4, 5})
	for i := 1; i <= 5; i++ {
		v, ok := iter.Next()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	for i := 0; i < 3; i++ {
		_, ok := iter.Next()
		assert.False(t, ok)
	}
}

func TestChanIterator(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	iter := itertools.OfChan(ch)
	for i := 1; i <= 5; i++ {
		v, ok := iter.Next()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	for i := 0; i < 3; i++ {
		_, ok := iter.Next()
		assert.False(t, ok)
	}
}

func TestMapIterator(t *testing.T) {
	iter := itertools.OfSlice([]int{1, 2, 3, 4, 5})
	iter = itertools.Map(iter, func(i int) int { return i * 2 })
	for i := 2; i <= 10; i += 2 {
		v, ok := iter.Next()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	for i := 0; i < 3; i++ {
		_, ok := iter.Next()
		assert.False(t, ok)
	}
}

func TestFlatMapIterator(t *testing.T) {
	iter := itertools.OfSlice([]int{1, 2, 3, 4, 5})
	iter = itertools.FlatMap(iter, func(i int) itertools.Iterator[int] {
		return itertools.Take(itertools.Repeat(i), i)
	})
	for i := 1; i <= 5; i++ {
		for j := 0; j < i; j++ {
			v, ok := iter.Next()
			assert.True(t, ok)
			assert.Equal(t, i, v)
		}
	}
	for i := 0; i < 3; i++ {
		_, ok := iter.Next()
		assert.False(t, ok)
	}
}

func TestFilterIterator(t *testing.T) {
	iter := itertools.OfSlice([]int{1, 2, 3, 4, 5})
	iter = itertools.Filter(iter, func(i int) bool { return i%2 == 0 })
	for i := 2; i <= 5; i += 2 {
		v, ok := iter.Next()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	for i := 0; i < 3; i++ {
		_, ok := iter.Next()
		assert.False(t, ok)
	}
}

func TestTakeIterator(t *testing.T) {
	iter := itertools.Take(itertools.Inf(), 3)
	for i := 1; i <= 3; i++ {
		v, ok := iter.Next()
		assert.True(t, ok)
		assert.Equal(t, i, v)
	}
	for i := 0; i < 3; i++ {
		_, ok := iter.Next()
		assert.False(t, ok)
	}
}
