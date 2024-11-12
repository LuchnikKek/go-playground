package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func clone(s string) string {
	b := make([]byte, len(s))
	copy(b, s)
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	// Проверка системных сигналов (эмуляция утечки памяти)
	// Раз в секунду строка копируется, без оптимизаций
	timeInNanosec := time.Now().UnixNano()
	rand.New(rand.NewSource(timeInNanosec))
	dataHolder := make([]string, 0)
	filler := randStringRunes(1024224)

	
	sigChannel := make(chan os.Signal, 1)
	// docker kill --signal=Код Контейнер
	// hangup (1) - отлавливается
	// interrupt (2) - отлавливается
	// kill (9) - exited(137)
	// docker run -d --memory=32m --memory-swap=32m myserv - exited(137)
	signal.Notify(sigChannel, os.Interrupt, os.Kill, syscall.SIGHUP)

	run := true

	go func() {
		oscall := <- sigChannel
		log.Printf("system call:%+v", oscall)
		run = false
		os.Exit(1)
	}()

	for {
		if !run {
			break
		}
		dataHolder = append(dataHolder, clone(filler))
		time.Sleep(1 * time.Second)
	}
}