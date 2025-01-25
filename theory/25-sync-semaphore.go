package theory

import (
	"fmt"
	"time"
)

func MainSyncSemaphore() {
	sem := NewSemaphore(3)
	for i := range 10 {
		sem.Acquire()
		go longFunc(i, sem)
	}
	fmt.Println("постановка (не выполнение) горутин запустит таймер")
	time.Sleep(10 * time.Second)
}

type Semaphore struct {
	cap int
	ch  chan int
}

func NewSemaphore(cap int) *Semaphore {
	ch := make(chan int, cap)
	s := &Semaphore{cap: cap, ch: ch}
	return s
}

func (s *Semaphore) Acquire() {
	s.ch <- 1
}

func (s *Semaphore) Release() {
	<-s.ch
}

func longFunc(i int, sem *Semaphore) {
	// sem.Acquire() - таймер выше будет запущен сразу, горутины не заблокируют
	defer sem.Release()
	fmt.Printf("started %v\n", i)
	time.Sleep(3 * time.Second)
	fmt.Printf("ended %v\n", i)
}
