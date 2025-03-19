package practice

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
)

type FloatDuration time.Duration

// Set записывает значение во FloatDuration
// Одновременно используется и для Flag, и для Env
func (d *FloatDuration) Set(val float64) {
	*d = FloatDuration(time.Duration(val * float64(time.Second)))
}

// String возвращает строковое представление time.Duration
func (d FloatDuration) String() string {
	return time.Duration(d).String()
}

// UnmarshalText вызывается при парсинге из ENV.
func (d *FloatDuration) UnmarshalText(text []byte) error {
	val, err := strconv.ParseFloat(string(text), 64)
	if err != nil {
		return err
	}
	d.Set(val)
	return nil
}

type EnvConfigTimeFloat struct {
	TaskDuration FloatDuration `env:"TASK_DURATION"`
	Timeout      FloatDuration `env:"TIMEOUT" envDefault:""`
}

func (ec *EnvConfigTimeFloat) Load() {
	taskDurFlag := flag.Float64("d", float64(ec.TaskDuration), "task duration")
	timeoutFlag := flag.Float64("t", float64(ec.Timeout), "timeout")
	flag.Parse()

	ec.TaskDuration.Set(*taskDurFlag)
	ec.Timeout.Set(*timeoutFlag)

	if err := env.Parse(ec); err != nil {
		log.Fatal(err)
	}
}

// TIMEOUT=3.5 go run main.go -d 2 -t 1.2
func MainConfigTimeFloatExample() {
	conf := EnvConfigTimeFloat{}
	conf.Load()
	log.Printf("loaded: %+v\r\n", conf)
}
