package theory

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Server struct {
	rqCount int64
}

const incrementCount int = 500

func mainRaceCondition() {
	// ПЛОХО, отсутствует Mutex, состояние гонки
	srv := &Server{}
	for range incrementCount {
		go increment(srv)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(srv.rqCount)
}

func increment(s *Server) {
	s.rqCount += 1
}
func mainWithMutex() {
	// ХОРОШО? но mutex излишний. Он нужен для блокировки выполнения кучи кода
	srv := &Server{}
	mu := &sync.Mutex{}
	for range incrementCount {
		go incrementMutex(srv, mu)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(srv.rqCount)
}

func incrementMutex(s *Server, mu *sync.Mutex) {
	mu.Lock()
	s.rqCount += 1
	mu.Unlock()
}

func mainWithAtomic() {
	// Отлично
	srv := &Server{}
	for range incrementCount {
		go incrementAtomic(srv)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(srv.rqCount)
}

func incrementAtomic(s *Server) {
	atomic.AddInt64(&s.rqCount, 1)
}

func MainSyncMutex() {
	mainRaceCondition()
	mainWithMutex()
	mainWithAtomic()
}
