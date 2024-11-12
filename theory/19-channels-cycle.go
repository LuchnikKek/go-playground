package theory

import "fmt"

func writeChan(ch chan<- int) {
	// chan<- = в функции можно только писать В канал
	// <-chan = в функции можно только читать ИЗ канала
	for i := 1; i <= 5; i++ {
		// блокируется, пока не прочитаем
		ch <- i
	}
	close(ch)
}

func MainChannelsCycle() {
	ch := make(chan int)

	go writeChan(ch)

	for v := range ch { // без close(ch) гарантирован дедлок
		fmt.Println(v)
	}
}


