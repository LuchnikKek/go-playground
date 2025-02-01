package theory

import (
	"fmt"
)

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

func MainPanicHandler() {
	err := safeFoo()
	if err != nil {
		fmt.Printf("Обработана ошибка: %v\n", err.Error())
	}
}

func safeFoo() (err error) {
	defer func() {
		if er := recover(); er != nil {
			err = fmt.Errorf("%v", er)
		}
	}()
	panic("Паника!!!")
}
