package theory

import (
	"log"
	"time"

	"github.com/caarlos0/env/v6"
)

// export TASK_DURATION=5s
// export FILES=test1.txt:test2.txt
// or
// TASK_DURATION=5s FILES=test1.txt:test2.txt go run main.go
func MainEnvCaarlos() {
	envStructType()
	envStructVar()
}

func envStructType() {
	type EnvConfig struct {
		Files []string `env:"FILES" envSeparator:":"`
		Home  string   `env:"HOME"`
		// required, чтобы переменная TASK_DURATION была определена
		TaskDuration time.Duration `env:"TASK_DURATION,required"`
	}

	var cfg EnvConfig

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg) // {[test1.txt test2.txt] /Users/ilakabesov 5s}
}

func envStructVar() {
	var userName struct {
		User string `env:"USER"`
	}

	if err := env.Parse(&userName); err != nil {
		log.Fatal(err)
	}

	log.Println(userName)
}
