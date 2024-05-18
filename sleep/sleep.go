package sleep

import (
	"runtime"
	"sync"
	"time"

	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Float | constraints.Integer](in []T) []T {
	if len(in) <= 1 {
		return in
	}

	ch := make(chan T, len(in))
	sema := make(chan struct{}, runtime.NumCPU())
	wg := sync.WaitGroup{}

	for i := 0; i < len(in); i++ {
		wg.Add(1)

		sema <- struct{}{}
		go func() {
			time.Sleep(acq * time.Duration(in[i]))
			ch <- in[i]

			<-sema
			wg.Done()
		}()
	}

	wg.Wait()
	close(ch)
	close(sema)

	var res = make([]T, 0, len(in))
	for v := range ch {
		res = append(res, v)
	}

	return res
}
