## What is Concurrency
**Concurrency** refers to programming language's ability to deal with lots of things at once.

**Parallelism** means doing lots of things simultaneously and independently. It might sound similar to concurrency, but it’s actually quite different.

## What is Goroutine
**Goroutine** is an independent function that executes simultaneously in some seperate lightweight threads managed by Go. (See: *Ref-1*)

## Race Conditions
**Race condition** occurs when two or more operations must execute in the correct order, but the program has not been written so that this order is guaranteed to be maintained.

## Memory Access Synchronization
Programs that modify data being simultaneously accessed by multiple goroutines must serialize such access. To serialize access, protect the data with channel operations or other synchronization primitives such as those in the sync and sync/atomic packages. (See: *Ref-3*)

## Deadlock
**Deadlocked** program is one in which all concurrent processes are waiting on one another. In this state, the program will never recover without outside intervention.

---
References:
1. https://go.dev/doc/effective_go#goroutines
2. https://go.dev/doc/articles/race_detector
3. https://go.dev/ref/mem