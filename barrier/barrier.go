package barrier

import (
	"fmt"
)

type Barrier struct {
	count    int // number of checkins
	total    int
	signalin chan bool
	done     chan bool
}

// NewBarrier creates a new barrier of size n.  will block n groutines unless call wait
func NewBarrier(n int) *Barrier {
	if n <= 0 {
		panic("Group must be >= 1")
	}

	b := &Barrier{0, n, make(chan bool, n), make(chan bool, n)}

	go controller(b)

	return b
}

func controller(b *Barrier) {

	for b.count != b.total {
		select {
		case <-b.signalin:
			b.count++
		}
	}

	// read from running channel to unlock

	for i := 0; i < b.total; i++ {

		b.done <- true
	}

}

func (b *Barrier) Wait() {
	b.signalin <- true

	if <-b.done {
		//fmt.Println(b.count)
		fmt.Println("max goroutines")

	}

}
