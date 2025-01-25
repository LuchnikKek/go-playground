package theory

import (
	"fmt"
	"runtime"
	"time"
)

func anyJob(num int) {
	fmt.Println(num)
	runtime.Gosched()
}

func MainGorutines() {
	fmt.Println(123)
	for i := 1; i <= 5; i++ {
		go anyJob(i)
	}
	time.Sleep(5 * time.Second)
}
