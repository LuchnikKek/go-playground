package theory

import "fmt"

func MainVariables() {
	// ========== Переменные ===========
	var v1 int // 0, zero values are 0, "", false
	var v2 int = 2 // 100
	v3 := 3 // var v3 int = 5
	fmt.Println(v1, v2, v3) // 0, 2, 3
	
	v1 = 10
	v2, v3 = 30, 20
	fmt.Println(v1, v2, v3) // 10, 30, 20

	v2, v3 = v3, v2 // =, т.к. переменные существуют
	fmt.Println(v1, v2, v3) // 10, 20, 30
	
	v1, v2, v3, v4 := 100, 200, 300, 400 // :=, т.к. v4 объявляется впервые
	fmt.Println(v1, v2, v3, v4)

	// блочное создание
	var (
		v01 = 1
		v02 = "string"
		v03 = false
		someIgnoredVariable = 12345
	)
	_ = someIgnoredVariable // игнорирование переменной
	fmt.Println("block:", v01, v02, v03)
	fmt.Println()

	// ========== Константы ===========
	const secondsInDay = 60 * 60 * 24
	fmt.Println("Seconds in day:", secondsInDay)

	const (
		minuteInSeconds = 60
		hourInMinutes = 60
		dayInHours = 24
	)
	fmt.Println("Seconds in day:", minuteInSeconds * hourInMinutes * dayInHours)
	fmt.Println()

	// ============ Указатели =============
	x := 123
	ptr := &x // получить указателя на адрес блока памяти
	fmt.Println(ptr, "points to", *ptr)

	xCopy := x // переприсваивание
	ptrxCopy := &xCopy
	fmt.Println(ptrxCopy)

	xCopyByPtr := *ptr // присваивание по значению указателя
	ptrxCopyByPtr := &xCopyByPtr
	fmt.Println(ptrxCopyByPtr)

	// ptr <> ptrxCopy <> ptrxCopyByPtr
}
