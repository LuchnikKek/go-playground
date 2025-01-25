package theory

import (
	"fmt"
	"time"
)

const (
	timerTime     = 3 * time.Second // время отвала по таймеру - Время вышло
	executionTime = 2 * time.Second // время имитации выполнения работы в джобе
	mainSleep     = 5 * time.Second // время, которое main ждёт. Обычный слип для ожидания горутины. Должен быть наибольшим
)

// Через executionTime проверяется таймер
// Если timerTime <= executionTime, 				то <-t.C. Даже quit() не поможет (выставлен if и редирект)
// Если timerTime > executionTime и quit <- 1 		то <-q
// Если timerTime > executionTime					то default

func MainSyncTimer() {
	timer := time.NewTimer(timerTime)
	quit := make(chan int)

	go jobWithTimeout(timer, quit)
	// go func() { quit <- 1 }()
	// quit <- 1 // нужна горутина - а то лок

	time.Sleep(mainSleep)
}

func jobWithTimeout(t *time.Timer, q chan int) {
	time.Sleep(executionTime)
	select {
	case <-t.C:
		fmt.Println("Время вышло")
	case <-q:
		if !t.Stop() { // if timer already expired/stopped
			<-t.C
		}
		fmt.Println("Таймер остановлен")
	default:
		fmt.Println("Завершение функции")
	}
}
