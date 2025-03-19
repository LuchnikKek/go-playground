package theory

import (
	"log"
	"time"

	"github.com/caarlos0/env/v6"
)

// TASK_DURATION=5s FILES=test1.txt:test2.txt go run main.go
func MainEnvCaarlosStruct() {
	conf := EnvConfig{}
	log.Printf("loaded: %+v\r\n", conf)
}

type EnvConfig struct {
	Files        []string      `env:"FILES" envSeparator:":"`
	Home         string        `env:"HOME"`
	TaskDuration time.Duration `env:"TASK_DURATION,required"`
}

func (ec *EnvConfig) Load() {
	if err := env.Parse(ec); err != nil {
		log.Fatal(err)
	}
}
