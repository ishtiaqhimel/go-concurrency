/*
Critical section of a program that needs exclusive access to a shared resource.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var data int

	// Here we add a variable that will allow our code to synchronize access to the data variable’s memory.
	var memoryAccess sync.Mutex

	go func() {
		// Here we declare that until we declare otherwise, our goroutine should have exclusive access to this memory.
		memoryAccess.Lock()
		data++ // Critical Section
		// Here we declare that the goroutine is done with this memory.
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if data == 0 { // Critical Section
		fmt.Println(data)
	} else {
		fmt.Println(data) // Critical Section
	}
	memoryAccess.Unlock()
}

/*
we have solved our data race, we haven’t actually solved our race condition!
The order of operations in this program is still nondeterministic;
we’ve just narrowed the scope of the nondeterminism a bit.

In this example, either the goroutine will execute first, or both our if and
else blocks will. We still don’t know which will occur first in any given
execution of this program.

Later, we’ll explore the tools to solve this kind of issue properly.

Synchronizing access to the memory in this manner also has performance ramifactions.
We’ll save the details for later when we examine the sync package, but the calls to Lock can
make our program slow. Every time we perform one of these operations, our program pauses for a period of time.
This brings up two questions:
	1. Are my critical sections entered and exited repeatedly?
	2. What size should my critical sections be?

Synchronizing access to the memory also shares some problems with other techniques of modeling concurrent problems -
	- Deadlocks
	- Livelocks
	- Starvation
*/
