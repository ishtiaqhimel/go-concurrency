/*
A data race happens when processes have to access the same variable concurrently i.e. one process reads from a memory location while another simultaneously writes to the exact same memory location.
A race condition is a flaw in a program regarding the timing/ordering of operations which disrupts the logic of the program and leads to erroneous results.
*/

package main

import (
	"fmt"
)

func main() {
	var data int
	go func() {
		data++
	}()

	// time.Sleep(1 * time.Second) // Bad solution

	if data == 0 {
		fmt.Println(data)
	}
}

/*
There are three possible outcomes to running this code:

- Nothing is printed. In this case, line 15 was executed before line 20.
- “0” is printed. In this case, lines 20 and 21 were executed before line 15.
- “1” is printed. In this case, line 20 was executed before line 15, but line 15 was executed before line 21.

// [golang official race-detector](https://go.dev/doc/articles/race_detector) holds some good examples of data races.
*/
