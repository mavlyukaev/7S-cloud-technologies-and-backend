package main

import "fmt"

type Mutex struct {
	Count  int
	signal chan struct{}
}

func (m *Mutex) Unlock() {
	m.signal <- struct{}{}
}

func (m *Mutex) Wait() {
	for i := 0; i < m.Count; i++ {
		<-m.signal
	}
}

func main() {
	m := Mutex{
		Count:  3,
		signal: make(chan struct{}, 3),
	}

	for i := 0; i < 3; i++ {
		go func() {
			defer m.Unlock()
			fmt.Printf("Горутина %d: Привет, мир!\n", i)
		}()
	}
	m.Wait()
}
