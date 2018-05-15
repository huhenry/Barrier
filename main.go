package main

import (
	"fmt"

	"github.com/huhenry/Barrier/barrier"
)

func main() {
	barrier := barrier.NewBarrier(10)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			barrier.Wait()
			fmt.Printf("unblock %d \n", i)
		}(i)
	}
}
