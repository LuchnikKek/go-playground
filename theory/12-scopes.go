package theory

import "fmt"

/*
Объекты и Поля с маленькой буквы видны только в пределах пакета.
Методы с маленькой буквы приватные

func someFunction() {...}
type user struct {...}


С заглавной буквы видны и за пределами пакета.
Методы с заглавной буквы публичные

func FunctionVisible() {...}
type ClientWithPrivate struct {
	VisibleValue int 		// будет видно, публичное поле
	hiddenValue string 		// не будет видно, приватное поле
}
*/

type ClientWithPrivate struct {
	VisibleValue int 		// будет видно
	hiddenValue string 		// не будет видно
}

// Конструктор. Позволит создавать объекты даже в других пакетах
func NewClientWithPrivate(visibleValue int, hiddenValue string) ClientWithPrivate {
	return ClientWithPrivate{
		VisibleValue: visibleValue,
		hiddenValue: hiddenValue,
	}
}

func (c *ClientWithPrivate) AddSuffixToHidden(suffix string) {
	c.hiddenValue += suffix
}

func MainScopes() {
	client := NewClientWithPrivate(5, "any/random/url")
	fmt.Println(client) 	// {5 any/random/url}
	client.AddSuffixToHidden("?userId=52")
	fmt.Println(client)		// {5 any/random/url?userId=52}
}
