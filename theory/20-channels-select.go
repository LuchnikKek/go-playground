package theory

import (
	"fmt"
	"runtime"
	"time"
)

func MainChannelsSelect() {
	// Мультиплексирование
	chanSelectBoth()
	fmt.Println("============LISTENER=============")
	chanListen()
}

func chanSelectBoth() {
	// пишет в оба канала. Отработает один рандомный
	// Оба канала не буферизированные, любое чтение/запись приведёт к блокировке
	ch := make(chan int)
	quit := make(chan int)

	// ch <- 1 // - deadlock

	go _chanWrite(ch)
	go _chanWrite(quit)

	runtime.Gosched()
	// выполняет горутины, порядок не определён

	select {
	case x := <-ch:
		// вычитаем значение из канала
		fmt.Println("ch =", x)
	case <-quit:
		// просто ждём, что в канале появится что-то, не суть
		fmt.Println("quit")
	}
}

func _chanWrite(ch chan<- int) {
	ch <- 1
}

func chanListen() {
	ch := make(chan int)
	quit := make(chan int)

	go _listener(ch, quit)
	go _chanWrite(ch)

	runtime.Gosched()

	go _chanWrite(quit)
	time.Sleep(1 * time.Second) // без sleep не успеет вывести "quit"
}

func _listener(ch, quit <-chan int) {
	// постоянно слушает канал
	// Как только что-то появляется в ch - делает
	// Как только что-то появляется в quit - прекращает цикл
	for {
		select {
		case x := <-ch:
			fmt.Println("ch = ", x)
		case <-quit:
			fmt.Println("quit")
			return
		default: // спасает от deadlock
			fmt.Println("default") // будет спамить, мб нужен просто слип
		}
	}
}
