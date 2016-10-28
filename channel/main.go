package main

import (
	"fmt"
	"sync"
)

type PongPongPayload struct {
	Counter int
}

func ExamplePingPong() {
	var p PongPongPayload
	chA := make(chan *PongPongPayload)
	chB := make(chan *PongPongPayload)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			p, ok := <-chA
			if !ok {
				break
			}
			fmt.Printf("chA: p.Counter = %d\n", p.Counter)
			p.Counter++
			if p.Counter > 6 {
				break

			}
			chB <- p
		}
		close(chB)
	}()

	go func() {
		defer wg.Done()
		for {
			p, ok := <-chB
			if !ok {
				break
			}
			fmt.Printf("chB: p.Counter = %d\n", p.Counter)
			p.Counter++
			if p.Counter > 6 {
				break

			}
			chA <- p
		}
		close(chA)
	}()
	chA <- &p
	wg.Wait()
}

func main() {
	ExamplePingPong()
}
