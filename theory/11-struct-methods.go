package theory

import "fmt"

func MainStructMethods() {
	MainStructMethodsByRef()
	MainStructMethodsByStruct()
}

// =========== Передача по ссылке ============
func MainStructMethodsByRef() {
	// user := UserWithPointer{
	// 	ID: 5,
	// 	Name: "Mike",
	// } // user.HasAvatar() == false

	// user := UserWithPointer{
	// 	ID: 5,
	// 	Name: "Mike",
	// 	Avatar: &IMG{},
	// } // user.HasAvatar() == false

	
	user := UserWithPointer{
		ID: 5,
		Name: "Mike",
		Avatar: &IMG{
			URL: "image_url",
			Size: 52,
		},
	}
	fmt.Println(user.Avatar.URL) // image_url
	user.ChangeAvatar()
	fmt.Println(user.Avatar.URL) // withRef
}

func (u UserWithPointer) HasAvatar() bool {
	return u.Avatar != nil && u.Avatar.URL != ""
}

func (u *UserWithPointer) ChangeAvatar() {
	u.Avatar.URL = "withRef"
}

// =========== Передача по значению ============
func MainStructMethodsByStruct() {
	user := User{
		ID: 5,
		Name: "Mike",
		Avatar: Image{
			URL: "image_url",
			Size: 52,
		},
	}
	fmt.Println(user.Avatar.URL) // image_url
	user.ChangeAvatar1()
	fmt.Println(user.Avatar.URL) // image_url
	user.ChangeAvatar2()
	fmt.Println(user.Avatar.URL) // withRef
}


func (u User) HasAvatar() bool {
	return u.Avatar.URL != ""
}

func (u User) ChangeAvatar1() {
	u.Avatar.URL = "withRef"
}

func (u *User) ChangeAvatar2() { // Жёсткая имба
	u.Avatar.URL = "withRef"
}
