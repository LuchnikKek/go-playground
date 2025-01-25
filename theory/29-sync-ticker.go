package theory

import (
	"fmt"
	"time"
)

func MainSyncTicker() {
	mainTicker()
	mainTickerInf()
}

func mainTicker() {
	// Можно остановить. Полезен для выполнения задач ограниченное число раз
	ticker := time.NewTicker(1 * time.Second)

	count := 0
	for tick := range ticker.C {
		count++
		fmt.Printf("Tick: #%v, time: %v\n", count, tick)
		if count > 5 {
			ticker.Stop()
			break
		}
	}
}

func mainTickerInf() {
	// Тикер, который нельзя остановить
	// Полезен для вечных фоновых задач, запускаемых вместе с приложением (типа мониторинга)
	ticker := time.Tick(1 * time.Second)

	count := 0
	for tick := range ticker {
		count++
		fmt.Printf("Tick: #%v, time: %v\n", count, tick)
	}
}
