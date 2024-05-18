package bogo

import (
	"cmp"
	"math/rand"
	"slices"
)

func Sort[T cmp.Ordered](in []T) []T {
	if len(in) <= 1 {
		return in
	}

	for !slices.IsSorted(in) {
		rand.Shuffle(len(in), func(i, j int) {
			in[i], in[j] = in[j], in[i]
		})
	}

	return in
}
