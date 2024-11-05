package theory

import "fmt"

type Sender interface {
	Send(msg string) error
}

type Email struct {
	Address string
}

func (e *Email) Send(msg string) error {
	fmt.Printf("Сообщение \"%v\" отправлено на адрес %v \n", msg, e.Address)
	return nil
}

func Notify(s Sender) {
	err := s.Send("Notify message")
	if err != nil {
		fmt.Println(err)
		return
	}

	// утверждение типов - type assertion
	// res, isPhone := s.(*Phone)
	// if isPhone {
	// 	fmt.Println("Success. Phone balance: ", res.Balance)
	// } else {
	// 	fmt.Println("Success")
	// }

	// type switch
	switch value := s.(type) {
	case *Phone:
		fmt.Println("Success. Phone balance: ", value.Balance)
	case *Email:
		fmt.Println("Success")
	}
}

type Phone struct {
	Number int
	Balance int
}

func (p *Phone) Send(msg string) error {
	fmt.Printf("Сообщение \"%v\" отправлено на номер %v \n", msg, p.Number)
	return nil
}

func MainInterfaces() {
	email := &Email{"ilya@gmail.com"}
	Notify(email)
	phone := &Phone{Number: 1323321, Balance: 100}
	Notify(phone) // полиморфизм	
}