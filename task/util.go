package task

import (
	"sync"
)

func merge(cs ...<-chan result) <-chan result {
	out := make(chan result)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c <-chan result) {
			defer wg.Done()

			for v := range c {
				out <- v
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func numsChan(start, end int) <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)

		for i := start; i <= end; i++ {
			c <- i
		}
	}()

	return c
}
