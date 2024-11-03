package theory

import "fmt"

// =========================================
// ======= Когда в структуре ссылка ========
type UserWithPointer struct {
	ID int64
	Name string
	Avatar *IMG
}

type IMG struct {
	URL string
	Size int64
}

func MainStructWithPointer() {
	user := UserWithPointer{}
	fmt.Printf("%#v\n", user) // theory.UserWithRef{ID:0, Name:"", Avatar:(*theory.IMG)(nil)}
	// IMG в этом случае - это ссылка. Никакой участок памяти ей не зарезервирован. Там nil
	// Все варианты ниже вызовут панику:
	// 		- user.Avatar.URL = "ssddsds"					= PANIC
	//		- user = _changeAvatar1(user, "withFullObj")	= PANIC
	//		- _changeAvatar2(&user, "withRef")				= PANIC
	//		- fmt.Println(user.Avatar.URL)					= PANIC

	// Чтобы избежать паники - нужно зарезервировать память:
	// user.Avatar = new(IMG) // либо так
	user.Avatar = &IMG{} // либо так

	// изменение значения. Вариант 1.
	user = _changeAvatar21(user)
	fmt.Println(user.Avatar.URL == "withFullObj")

	// изменение значения. Вариант 2.
	_changeAvatar22(&user)
	fmt.Println(user.Avatar.URL == "withRef")
}

func _changeAvatar21(user UserWithPointer) UserWithPointer {
	user.Avatar.URL = "withFullObj"
	return user
}

func _changeAvatar22(user *UserWithPointer) {
	user.Avatar.URL = "withRef"
}
