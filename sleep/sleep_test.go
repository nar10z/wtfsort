package sleep

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestSort(t *testing.T) {
	t.Parallel()

	t.Run("#1. Empty", func(t *testing.T) {
		t.Parallel()

		var in []int
		out := Sort(in)

		assert.Empty(t, out)
	})
	t.Run("#2. Single", func(t *testing.T) {
		t.Parallel()

		in := []int{1}
		out := Sort(in)

		assert.Equal(t, in, out)
	})
	t.Run("#3. Sorted input", func(t *testing.T) {
		t.Parallel()

		in := []int{1, 2, 3}
		out := Sort(in)

		assert.GreaterOrEqual(t, len(in), len(out))
		assert.True(t, slices.IsSorted(out))
	})
	t.Run("#4. Unsorted input", func(t *testing.T) {
		t.Parallel()

		in := []int{4, 3, 1, 2, 5, 6}
		out := Sort(in)

		fmt.Println(out)
		assert.GreaterOrEqual(t, len(in), len(out))
		assert.True(t, slices.IsSorted(out))
	})
}
