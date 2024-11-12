package theory

import "fmt"

func MainChannels() {
	unbufferedChannel()
	bufferedChannel()
}

func unbufferedChannel() {
	// небуферизированный канал (нет ёмкости)
	var ch chan int // объявление канала
	ch = make(chan int)

	// ch <- 100 - deadlock, канал ещё никто не читает

	go _readChan(ch)

	// ch <- 52 - если закомментировать = утечка горутин.
	// горутина будет висеть в фоне, пока main работает
	ch <- 52
}

func _readChan(ch chan int) {
	value := <-ch
	fmt.Println(value)
}

func bufferedChannel() {
	var ch chan int // объявление канала
	ch = make(chan int, 1) // создание буферизированного канала (защита от deadlock)

	ch <- 100

	go _readChan(ch) // 100

	ch <- 200 // не deadlock
}
