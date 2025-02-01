package theory

import (
	"flag"
	"log"
)

// Флаг запуска относится к самой джава-машине,
// а не является атрибутом бизнес-приложения как переменные окружения
// например для Java: кол-во памяти у Java-машины
// Ещё он принадлежит конкретному (!) запуску

// Флаги = консольные утилиты/системные настройки (память)
// env = пароли, большие данные

func MainStartupFlags() {
	// -help выдаст значения
	// --word		a string (foo as default)
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 52, "an int")     // --numb=22 or -numb 22
	boolPtr := flag.Bool("fork", false, "a bool") // --fork/-fork = true

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var") // ["lol", "kek"]

	flag.Parse()

	log.Printf("%#v", *wordPtr)
	log.Printf("%#v", *numbPtr)
	log.Printf("%#v", *boolPtr)
	log.Printf("%#v", svar)
	log.Printf("args: %v", flag.Args())
}
