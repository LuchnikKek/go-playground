package practice

import (
	"flag"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
)

type EnvConfigTimeDur struct {
	TaskDuration time.Duration `env:"TASK_DURATION"`
	Timeout      time.Duration `env:"TIMEOUT"`
}

func (ec *EnvConfigTimeDur) Load() {
	ec.loadFlags()
	ec.loadEnv() // overrides
}

func (ec *EnvConfigTimeDur) loadFlags() {
	flag.DurationVar(&ec.TaskDuration, "d", ec.TaskDuration, "task duration")
	flag.DurationVar(&ec.Timeout, "t", ec.Timeout, "timeout")

	flag.Parse()
}

func (ec *EnvConfigTimeDur) loadEnv() {
	if err := env.Parse(ec); err != nil {
		log.Fatal(err)
	}
}

// TIMEOUT=3 go run main.go -d 2 -t 1.2
func MainConfigTimeDurExample() {
	conf := EnvConfigTimeDur{}
	conf.Load()
	log.Printf("loaded: %+v\r\n", conf)
}
