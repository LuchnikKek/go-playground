package theory

import (
	"context"
	"fmt"
	"time"
)

func MainSyncContext() {
	mainCtxWithCancel()
	mainCtxWithTimeout()
}

func mainCtxWithCancel() {
	fmt.Println("Context with cancel")
	ctx, cancel := context.WithCancel(context.Background())
	for i := range 4 {
		go sendData(ctx, i)
	}
	time.Sleep(time.Second)
	cancel()
	time.Sleep(500 * time.Millisecond)
}

func sendData(ctx context.Context, num int) {
	timer := time.NewTimer(time.Duration(num) * time.Second)

	select {
	case <-ctx.Done():
		fmt.Printf("Sending #%v: canceled\n", num)
		return
	case <-timer.C:
		fmt.Printf("Sending #%v: success\n", num)
	}
}

func mainCtxWithTimeout() {
	fmt.Println("Context with timeout")
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 2*time.Second) // через сколько всё полетит в timer
	defer cancel()

	for i := range 4 {
		go sendData(ctx, i)
	}
	time.Sleep(3 * time.Second)
}
