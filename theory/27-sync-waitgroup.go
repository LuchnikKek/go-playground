package theory

import (
	"fmt"
	"sync"
	"time"
)

func sleep(t time.Duration, wg *sync.WaitGroup) {
	fmt.Println("sleep", t)
	time.Sleep(t)
	wg.Done()
}

func MainSyncWaitGroup() {
	// Wait() ждёт, пока счётчик не станет равен 0
	wg := &sync.WaitGroup{}

	for i := range 4 {
		wg.Add(1)
		duration := time.Duration(i) * time.Second
		go sleep(duration, wg)
	}
	wg.Wait()
}
