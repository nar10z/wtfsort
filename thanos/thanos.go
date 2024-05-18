package thanos

import (
	"cmp"
	"math/rand"
	"slices"
)

func Sort[T cmp.Ordered](in []T) []T {
	for !slices.IsSorted(in) {
		var (
			l   = len(in)
			res = make([]T, 0, l)
		)

		if l <= 1 {
			return in
		}

		// For example, if there are 7 living beings in the world, Thanos can destroy either 3 or 4 of them.
		// However, the comics and films do not specify how the Glove decides such mathematical nuances.
		// One can assume that it was left up to Thanos himself or the Infinity Glove to decide
		// on the exact number of creatures to destroy.
		//
		// We're not that violent, so we'll keep one element alive.
		// 3/2 = float64(1.5) -> int(1)
		// 4/2 = float64(2) -> int(2)
		qntyDead := l / 2
		uniqueIdx := make(map[int]struct{}, qntyDead)
		for i := 0; i < qntyDead; i++ {
			for {
				rnd := rand.Intn(l)
				if _, ok := uniqueIdx[rnd]; ok {
					continue
				}

				uniqueIdx[rnd] = struct{}{}

				break
			}
		}

		for i := range in {
			if _, ok := uniqueIdx[i]; ok {
				continue
			}

			res = append(res, in[i])
		}

		in = res
	}

	return in
}
