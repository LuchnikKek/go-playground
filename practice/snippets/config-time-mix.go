package practice

import (
	"flag"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
)

type EnvConfigTime struct {
	TaskDuration time.Duration `env:"TASK_DURATION"`
	Timeout      time.Duration `env:"TIMEOUT"`
}

func (ec *EnvConfigTime) Load() {
	ec.loadFlags()
	ec.loadEnv() // overrides
}

func (ec *EnvConfigTime) loadFlags() {
	taskDur := flag.Float64("d", ec.TaskDuration.Seconds(), "task duration")
	timeout := flag.Float64("t", ec.Timeout.Seconds(), "timeout")

	flag.Parse()

	ec.TaskDuration = time.Duration(*taskDur * float64(time.Second))
	ec.Timeout = time.Duration(*timeout * float64(time.Second))
}

func (ec *EnvConfigTime) loadEnv() {
	if err := env.Parse(ec); err != nil {
		log.Fatal(err)
	}
}

// TIMEOUT=3s go run main.go -d 2 -t 1.2
func MainConfigTimeExample() {
	conf := EnvConfigTime{}
	conf.Load()
	log.Printf("loaded: %+v\r\n", conf)
}
