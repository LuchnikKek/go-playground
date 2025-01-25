package theory

import "fmt"

func foo() {
	panic("Паника в foo()")
}

func MainPanic() {
	defer func() { // не работает с паникой из горутины, т.к. у горутины свой стек
		err := recover()
		fmt.Println(err)
		fmt.Println("Восстановление после паники")
	}()
	fmt.Println("Старт")
	foo()
	fmt.Println("Финиш") // не выведется
}
