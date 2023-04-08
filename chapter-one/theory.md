## What is Concurrency
**Concurrency** refers to programming language's ability to deal with lots of things at once.

**Parallelism** means doing lots of things simultaneously and independently. It might sound similar to concurrency, but it’s actually quite different.

## What is Goroutine
**Goroutine** is an independent function that executes simultaneously in some seperate lightweight threads managed by Go. (See: *goroutine.go* & *Ref-1*)

## Race Conditions
**Race condition** occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained.

---
References:
1. https://go.dev/doc/effective_go#goroutines
2. https://go.dev/doc/articles/race_detector