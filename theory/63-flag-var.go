package theory

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// Реализует интерфейс flag.Value -
type NetAddress struct {
	Host string
	Port int
}

// Возвращает строковое значение флага
func (n *NetAddress) String() string {
	return fmt.Sprint(n.Host, ":", n.Port)
}

// Принимает значение флага и разбирает в структуру
func (n *NetAddress) Set(flagValue string) error {
	split := strings.Split(flagValue, ":")
	if len(split) != 2 {
		return fmt.Errorf("invalid address format: %s", flagValue)
	}

	port, err := strconv.Atoi(split[1])
	if err != nil {
		return err
	}

	n.Host, n.Port = split[0], port
	return nil
}

func MainStartupFlagsVar() {
	addr := new(NetAddress)

	// бесполезно, просто проверка на соответствие интерфейсу
	_ = flag.Value(addr)

	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()

	fmt.Println(addr.Host)
	fmt.Println(addr.Port)
}
