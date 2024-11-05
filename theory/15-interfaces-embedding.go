package theory

import "fmt"

type Caller interface {
	Call(number int) error
}

type SmartPhone interface {
	Sender
	Caller
	OneMoreMethod()
}

// Структура, соответствующая интерфейсу SmartPhone
type Xiaomi struct {}

func(x *Xiaomi) Send(msg string) error {
	fmt.Println("stub")
	return nil
}

func(x *Xiaomi) Caller(msg string) error {
	fmt.Println("stub")
	return nil
}

func(x *Xiaomi) OneMoreMethod() {
	fmt.Println("stub")
}
