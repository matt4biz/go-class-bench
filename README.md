[![Run on Repl.it](https://repl.it/badge/github/matt4biz/go-class-bench)](https://repl.it/github/matt4biz/go-class-bench)

# Go class: Benchmark examples
Note that the results may vary depending on the version of Go, the operating system, the underlying hardware, and perhaps the phase of the moon. Benchmarking is a black art. However, these examples should show quite a difference in almost any case, simple as they are.

## Fibonacci
This example runs two versions of Fibonacci; one recursive, one not. 

Notice that while the recursive version calls itself almost 22000 times, it still costs 4-5 times as much to do each addition when we wrap it with function calls (the function is not dynamically dispatched and runs entirely from L1 cache, i.e., this is as good as it gets).

`go test -bench=. fib/fib_test.go`

## List v Slice
This example shows the difference between a slice and a singly linked list.

### List with value in the node
The first version measures the cost to make and iterate over the list or slice.

`go test -bench=. list1/list1_test.go`

The second builds the list or slice before the loop and resets the 
benchmark timer.

`go test -bench=. list2/list2_test.go`

### List with value held separately
Run this one benchmarking memory allocations; the loop both allocates and iterates (as with the first example). 

`go test -bench=. -benchmem list3/list3_test.go`

## False Sharing
The first version exhibits false sharing because all the counters are in the same cache line (we allocated them deliberately this way) and are updated for each value read from the channel.

`go test -bench=. -benchtime=10s -cpu=2,4,8 ./share1/share1_test.go`

The second version has much better performance when we eliminate that false sharing by keeping a local counter and writing back once at the end.

`go test -bench=. -benchtime=10s -cpu=2,4,8 ./share2/share2_test.go`

Note that there's probably some other contention in this program due to the channels; it's a bit artificial, but the impact of false sharing is very clear.