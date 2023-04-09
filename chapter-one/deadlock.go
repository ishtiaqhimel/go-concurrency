package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func deadlock() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()         // attempt to enter critical section
		defer v1.mu.Unlock() // exit the critical section before printSum returns

		time.Sleep(2 * time.Second) // here we sleep for a period of time to simulate work (and trigger a deadlock)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(1)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}

func main() {
	deadlock()
}

/*
The Coffman Conditions ( that help detect, prevent, and correct deadlocks) are as follows:
1. Mutual Exclusion
	- A concurrent process holds exclusive rights to a resource at any one time.
2. Wait For Condition
	- A concurrent process must simultaneously hold a resource and be waiting for an additional resource.
3. No Preemption
	- A resource held by a concurrent process can only be released by that process, so it fulfills this condition.
4. Circular Wait
	- A concurrent process (P1) must be waiting on a chain of other concurrent processes (P2), which are in turn
	waiting on it (P1), so it fulfills this final condition too.

These laws allow us to prevent deadlocks too. If we ensure that at least one of these conditions is not true,
we can prevent deadlocks from occurring. Unfortunately, in practice these conditions can be hard to reason about,
and therefore difficult to prevent.
*/
