package itertools_test

import (
	"testing"

	"github.com/qsliu2017/go-itertools"
	"github.com/stretchr/testify/assert"
)

func TestToChan(t *testing.T) {
	iter := itertools.OfSlice([]int{1, 2, 3, 4, 5})
	ch := itertools.ToChan(iter)
	for i := 1; i <= 5; i++ {
		assert.Equal(t, i, <-ch)
	}
	for i := 0; i < 3; i++ {
		_, ok := <-ch
		assert.False(t, ok)
	}
}

func TestReduce(t *testing.T) {
	iter := itertools.OfSlice([]int{1, 2, 3, 4, 5})
	sum := itertools.Reduce(iter, 0, func(a, b int) int { return a + b })
	assert.Equal(t, 15, sum)
}

func TestToSlice(t *testing.T) {
	iter := itertools.OfSlice([]int{1, 2, 3, 4, 5})
	slice := itertools.ToSlice(iter)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, slice)
}

func TestGroupBy(t *testing.T) {
	iter := itertools.OfSlice([]int{1, 2, 3, 4, 5})
	group := itertools.GroupBy(iter, func(i int) int { return i % 2 }, func(i int) int { return i })
	assert.Equal(t, map[int]itertools.Set[int]{
		0: {2: {}, 4: {}},
		1: {1: {}, 3: {}, 5: {}},
	}, group)
}
