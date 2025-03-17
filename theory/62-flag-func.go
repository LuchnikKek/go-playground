package theory

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

// go run main.go -effects foo,bar,buz
func MainStartupFlagsFunc() {
	_ = parseArgsFunc
	parseArgsFuncVar()
}

func parseArgsFunc() {
	var effects []string

	// если встречается флаг - запускается обработчик
	flag.Func("effects", "Rotation and mirror", func(flagValue string) error {
		effects = strings.Split(flagValue, ",")
		return nil
	})

	flag.Parse()

	log.Println(effects)
}

// То же самое, но через flag.Var() и структуру для типизации
// см. flag.Value
type options struct {
	thumb   bool
	effects []string
}

// String() должен уметь сериализовать переменную типа в строку
func (o *options) String() string {
	return strings.Join(o.effects, ",")
}

func (o *options) Set(flagValue string) error {
	o.effects = strings.Split(flagValue, ",")
	return nil
}

func parseArgsFuncVar() {
	opts := new(options)

	flag.Var(opts, "effects", "Rotation and mirror")
	flag.BoolVar(&opts.thumb, "thumb", false, "create thumb") // без этого не видит thumb

	flag.Parse()

	fmt.Println(opts.effects)
	fmt.Println(opts.thumb)
}
