package main

import (
	"fmt"
	"sync"
	"time"
)

var dumplings = 20
var dumplingMutex sync.Mutex

type Chopstick struct {
	sync.Mutex
	name string
}

type Philosopher struct {
	name           string
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func (p Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		// Lock left chopstick
		p.leftChopstick.Lock()
		fmt.Printf("%s grabbed by %s, now needs %s\n", p.leftChopstick.name, p.name, p.rightChopstick.name)

		// Lock right chopstick
		p.rightChopstick.Lock()
		fmt.Printf("%s grabbed by %s\n", p.rightChopstick.name, p.name)

		// Critical section
		dumplingMutex.Lock()
		if dumplings > 0 {
			dumplings--
			fmt.Printf("%s eats a dumpling. Dumplings left: %d\n", p.name, dumplings)
		}
		dumplingMutex.Unlock()

		// Release right and left chopsticks
		p.rightChopstick.Unlock()
		fmt.Printf("%s released by %s\n", p.rightChopstick.name, p.name)

		p.leftChopstick.Unlock()
		fmt.Printf("%s released by %s\n", p.leftChopstick.name, p.name)

		fmt.Printf("%s is thinking...\n", p.name)
		time.Sleep(100 * time.Millisecond)

		// Stop if dumplings are gone
		dumplingMutex.Lock()
		if dumplings <= 0 {
			dumplingMutex.Unlock()
			break
		}
		dumplingMutex.Unlock()
	}
}

func main() {
	// Create chopsticks
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = &Chopstick{name: fmt.Sprintf("Chopstick-%d", i)}
	}

	// Create philosophers
	philosophers := make([]Philosopher, 5)
	for i := 0; i < 5; i++ {
		left := chopsticks[i]
		right := chopsticks[(i+1)%5]
		philosophers[i] = Philosopher{
			name:           fmt.Sprintf("Philosopher-%d", i),
			leftChopstick:  left,
			rightChopstick: right,
		}
	}

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go philosophers[i].dine(&wg)
	}

	wg.Wait()
	fmt.Println("All dumplings eaten.")
}
