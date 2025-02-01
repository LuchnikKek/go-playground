package theory

import (
	"fmt"
	"log"
	"os"

	"github.com/qiangxue/go-env"
)

func MainOsEnvs() {
	envBasics()

	envMapperShort()
	envMapperLong()
}

func envBasics() {
	_ = os.Setenv("FOO", "1") // set

	// get
	log.Printf("%q\n", os.Getenv("FOO")) // "1"
	log.Printf("%q\n", os.Getenv("BAR")) // ""

	_ = os.Setenv("BAR", "")

	_checkIfSet("BAR") // BAR is ""
	_checkIfSet("BUZ") // BUZ was never set

	// for _, e := range os.Environ() {
	// 	log.Println(e) // "KEY=VALUE"
	// }
}

func _checkIfSet(key string) {
	// os.LookupEnv(key) -> value, true if was set
	if val, isSet := os.LookupEnv(key); isSet {
		log.Printf("%v is %q", key, val)
	} else {
		log.Printf("%v was never set", key)
	}
}

func envMapperShort() {
	type configEnv struct {
		Host string
		Port int
	}

	_ = os.Setenv("APP_HOST", "localhost")
	_ = os.Setenv("APP_PORT", "5432")

	var cfg configEnv

	if err := env.Load(&cfg); err != nil {
		panic(err)
	}

	fmt.Println(cfg.Host) // "localhost"
	fmt.Println(cfg.Port) // 5432
}

func envMapperLong() {
	type configEnv struct {
		Host string `env:"PLAYGROUND_HOST"`
		Port int    `env:"PLAYGROUND_PORT"`
	}

	_ = os.Setenv("APP_PLAYGROUND_HOST", "127.0.0.1")
	_ = os.Setenv("APP_PLAYGROUND_PORT", "9092")

	var cfg configEnv

	loader := env.New("APP_", log.Printf)
	if err := loader.Load(&cfg); err != nil {
		panic(err)
	}

	fmt.Println(cfg.Host) // "127.0.0.1"
	fmt.Println(cfg.Port) // 9092
}
