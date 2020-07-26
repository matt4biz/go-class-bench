package share1

import (
	"sync"
	"testing"
)

const (
	nchan   = 8
	nworker = nchan
	buffer  = 1024
)

var wg sync.WaitGroup

func count(cnt *uint64, in <-chan int) {
	for i := range in {
		*cnt += uint64(i)
	}

	wg.Done()
}

func fill(n int, in chan<- int) {
	for i := 0; i < n; i++ {
		in <- i
	}
	close(in)
}

func run() (total int) {
	var cnt [nworker]uint64

	in := make([]chan int, nchan)

	for i := 0; i < nchan; i++ {
		in[i] = make(chan int, buffer)
		go fill(10000, in[i])
	}

	wg.Add(nworker)

	for i := 0; i < nworker; i++ {
		go count(&cnt[i], in[i%nchan])
	}

	wg.Wait()

	for _, v := range cnt {
		total += int(v)
	}
	return
}

func BenchmarkShare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		run()
	}
}
