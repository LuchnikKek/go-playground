package theory

import "fmt"

// =========================================
// ======= Когда в структуре структура ========

type User struct {
	ID int64
	Name string
	Avatar Image
}

type Image struct {
	URL string
	Size int64
}

func MainStructWithStruct() {
	user := User{}
	fmt.Printf("%#v\n", user) // theory.User{ID:0, Name:"", Avatar:theory.Image{URL:"", Size:0}}
	
	// изменение значения. Вариант 1.
	user = _changeAvatar1(user)
	fmt.Println(user.Avatar.URL == "withFullObj")

	// изменение значения. Вариант 2.
	_changeAvatar2(&user) // "withRef"
	fmt.Println(user.Avatar.URL == "withRef")
}

func _changeAvatar1(user User) User {
	user.Avatar.URL = "withFullObj"
	return user
}

func _changeAvatar2(user *User) {
	user.Avatar.URL = "withRef"
}
