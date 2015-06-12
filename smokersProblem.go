

package main

import (
	"fmt"
	"math/rand"
	"sync"
)
// Implementation of a Barrier from Book with ISBN: 978-2-642-29968-1/Page 103
type Barrier struct {
	M, n     uint
	mutex, s sync.Mutex
}

func createBarrier(n uint) *Barrier {
	x := new(Barrier)
	x.M = n
	x.s.Lock()
	return x
}
func (x *Barrier) Wait() {
	x.mutex.Lock()
	x.n++
	if x.n == x.M {
		if x.n == 1 {
			x.mutex.Unlock()
		} else {
			x.n--
			x.s.Unlock()
		}
	} else {
		x.mutex.Unlock()
		x.s.Lock()
		x.n--
		if x.n == 0 {
			x.mutex.Unlock()
		} else {
			x.s.Unlock()
		}
	}
}
// End of implementation of Barrier

var (
	avail                [3]sync.Mutex
	ready, locked, mutex sync.Mutex
	done                 chan bool = make(chan bool)
	round                uint
	barrier             *Barrier            
)

func agent(u uint) {
	ready.Lock()
	avail[(u+1)%3].Unlock()
	avail[(u+2)%3].Unlock()
}
func Wirtin() {
	smokerOut()
	for {
		u := uint(rand.Intn(3))
		agent(u)
		fmt.Printf("Waiter has put %d and %d on the table\n", (u+1)%3, (u+2)%3)
	}
}

func smokerIn(u uint) {
	mutex.Lock()
	fmt.Printf("Ich am smoker %d \n", u)
	avail[(u+1)%3].Lock()
	avail[(u+2)%3].Lock()
	mutex.Unlock()
}

func smokerOut() { 
	ready.Unlock()
}

func test(u uint) {
	// Smoker u has already one of the things on the table
	if avail[u] != locked {
		barrier.Wait()
		Raucher(u)
	// Nothing is on the table
	}else if avail[(u+1)%3] == locked && avail[(u+2)%3]== locked{
		barrier.Wait()
		Raucher(u)
	}
}
/*test(u) makes sure that only the smoker who does not have the things
layen on the table is also the only one who can smoke. The others wait
at the barrier. Wenn the smokers who smokes is done, he goes to also to 
the Barrier, and that way the 3 are ready again at the same time.  */
func Raucher(u uint) {
	for {
		test(u)
		smokerIn(u)
		fmt.Printf("Smoker %d smokes\n", u)
		if round+1 < 10 {
			round++
		} else {
			done <- true
			return
		}
		fmt.Printf("Smoker %d smokes no more\n", u)
		smokerOut()
		barrier.Wait()
	}
}
func main() {
	ready.Lock()
	locked.Lock()
	barrier = createBarrier(3)
	for u := uint(0); u < 3; u++ {
		avail[u].Lock()
		raucher[u].Lock()
	}
	go Wirtin()
	for u := uint(0); u < 3; u++ {
		go Raucher(u)
	}
	<-done
}
