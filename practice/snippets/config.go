package practice

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
)

type EnvConfig struct {
	Files        []string      `env:"FILES" envSeparator:":"`
	Home         string        `env:"HOME"`
	TaskDuration time.Duration `env:"TASK_DURATION"`
}

func (ec *EnvConfig) Load() {
	ec.loadFlags()
	ec.loadEnv() // overrides flags
}

func (ec *EnvConfig) loadFlags() {
	flag.Func("f", "input files", func(flagValue string) error {
		ec.Files = strings.Split(flagValue, ":")
		return nil
	})
	flag.StringVar(&ec.Home, "home", ec.Home, "home dir path")
	flag.DurationVar(&ec.TaskDuration, "d", ec.TaskDuration, "task duration")
	flag.Parse()
}

func (ec *EnvConfig) loadEnv() {
	if err := env.Parse(ec); err != nil {
		log.Fatal(err)
	}
}

// FILES=test1.txt:test2.txt go run main.go -d 1s
func MainConfigExample() {
	conf := EnvConfig{}
	conf.Load()
	log.Printf("loaded: %+v\r\n", conf)
}
