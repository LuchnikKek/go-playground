package theory

import (
	"errors"
	"fmt"
)

// Relationship определяет положение в семье.
type Relationship string

// Возможные роли в семье.
const (
	Father      = Relationship("father")
	Mother      = Relationship("mother")
	Child       = Relationship("child")
	GrandMother = Relationship("grandMother")
	GrandFather = Relationship("grandFather")
)

// Family описывает семью.
type Family struct {
	Members map[Relationship]Person
}

// Person описывает конкретного человека в семье.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// ErrRelationshipAlreadyExists возвращает ошибку, если роль уже занята.
var ErrRelationshipAlreadyExists = errors.New("relationship already exists")

// AddNew добавляет нового члена семьи.
// Если в [Family] ещё нет [Person], создаётся пустая мапа.
// Если роль уже занята, метод выдаёт [ErrRelationshipAlreadyExists].
func (f *Family) AddNew(r Relationship, p Person) error {
	if f.Members == nil {
		f.Members = map[Relationship]Person{}
	}
	if _, ok := f.Members[r]; ok {
		return ErrRelationshipAlreadyExists
	}
	f.Members[r] = p
	return nil
}

// FullName возвращает Имя + Фамилию
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func MainRelationship() {
	f := Family{}
	err := f.AddNew(Father, Person{
		FirstName: "Misha",
		LastName:  "Popov",
		Age:       56,
	})
	fmt.Println(f, err)

	err = f.AddNew(Father, Person{
		FirstName: "Drug",
		LastName:  "Mishi",
		Age:       57,
	})
	fmt.Println(f, err)
}
