package theory

import "fmt"

func NotifyEmpty(i interface{}) {
	// switch на конкретные типы
	switch i.(type) {
	case int:
		fmt.Println("Тип int не поддерживается")
	}

	// утверждение конкретного интерфейса
	s, ok := i.(Sender)
	if !ok {
		fmt.Println("Ошибка утверждения интерфейса")
		return
	}

	err := s.Send("Сообщение в пустой интерфейс")
	if err != nil {
		fmt.Println("Произошла ошибка")
		return
	}
	fmt.Println("Success")
}

func MainInterfacesEmpty() {
	var b interface{}
	fmt.Println(b) // nil, т.к. указан тип как пустой интерфейс, но не объявлено значение

	NotifyEmpty(2)
	email := &Email{"ilya@gmail.com"}
	NotifyEmpty(email)
	NotifyEmpty("lol")
	sl := [5]int64{1, 2, 3, 4, 5}
	NotifyEmpty(sl)
}
