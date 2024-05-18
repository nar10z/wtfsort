package stalin

import (
	"cmp"
)

func Sort[T cmp.Ordered](in []T) []T {
	if len(in) <= 1 {
		return in
	}

	res := []T{in[0]}

	for i := 1; i < len(in); i++ {
		if res[len(res)-1] > in[i] {
			continue
		}

		res = append(res, in[i])
	}

	return res
}
