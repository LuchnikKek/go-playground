package theory

import (
	"fmt"
	"strings"
)

func mainFunctions() {
	mainFuncs()
	mainRecursion()
	mainLambdas()
	mainClosures()
	mainDefer()
}

func mainFuncs() {
	// =========== Функции ============
	fmt.Println("======= Функции =========")

	fmt.Println(_returnTyped() == _returnNamed()) // true
	
	fmt.Println(_variativeFuncArray(1, []int{2, 3, 4})) // [2 3 4 1]
	fmt.Println(_variativeFunc(1, 2, 3, 4)) // [2 3 4 1]
	fmt.Println(_variativeFunc(1, []int{2, 3, 4}...)) // распаковка слайса при передаче! [2 3 4 1]

	newStr, err := _addSuffixWithErr("someStr")
	fmt.Println(newStr) // someStrSuffix
	fmt.Println(err) // <nil>
}

func _returnTyped() int {
	// результат по типу
	lol := 123
	return lol
}

func _returnNamed() (lol int) {
	// именованное возвращение результата
	lol = 123
	return
}

func _variativeFuncArray(item int, listSlice []int) (newSlice []int) {
	// принимает элемент и список
	newSlice = append(listSlice, item)
	return
}

func _variativeFunc(item int, listSlice ...int) (newSlice []int) {
	// принимает элементы, запаковывает их
	newSlice = append(listSlice, item)
	return
}

func _addSuffixWithErr(origin string) (string, error) {
	// возвращает строку с суффиксом и отсутствие ошибок
	newOrigin := origin + "Suffix"
	return newOrigin, nil
}

func mainRecursion() {
	// Рекурсия
	fmt.Println("=========== Рекурсия ==============")
	recursionResult := _recursionFactorial(5) // рекурсия для получения факториала
	fmt.Println(recursionResult)
}

func _recursionFactorial(n uint) uint {
	// функция по нахождению факториала числа через рекурсию
	if n <= 1 {
		return 1
	}
	return n * _recursionFactorial(n - 1)
}

func mainLambdas() {
	fmt.Println("=========== Анонимные функции (lambda) ==============")
	// анонимные функции f := func () { ... }
	// func(r rune) rune { return r + 1 }
	lambdaResult := strings.Map(func(r rune) rune { return r + 1 }, "SDWS")
	fmt.Println(lambdaResult)

	// присваивание анонимной функции (моветон, по идее?)
	var f1, f2 func(s string) int

	f1 = func(s string) int { return len(s) }
	f2 = func(s string) int { return len(s) + 50 }

	fmt.Println(f1("ss"), f2("ss")) // 2, 52
}

func mainClosures() {
	// Замыкания
	fmt.Println("=========== Замыкания (closure) ==============")
	// замыкание (closure) без передачи значений
	closed_func_1 := _counterClosure()
	closed_func_1()
	closed_func_1()

	// замыкание (closure) с передачей значений
	closedSumFunc1 := _parametrizedSumClosure(1)
	closedSumFunc1(2)
	resSum1 := closedSumFunc1(3)
	fmt.Println("first:", resSum1) // first: 6

	closedSumFunc2 := _parametrizedSumClosure(10)
	closedSumFunc2(15)
	resSum2 := closedSumFunc2(16)
	fmt.Println("second:", resSum2) // second: 41
}

func _counterClosure() func() {
	counter := 0
	return func () {
		counter += 1
		fmt.Println("Печатаю", counter)
	}
}

func _parametrizedSumClosure(start int) func(int) int {
	return func (summand int) int {
		start += summand
		return start
	}
}

func mainDefer () {
	// Defer выполняется при закрытии функции
	fmt.Println("=========== Отложенный вызов (Defer) ==============")
	// если defer несколько - пойдут с конца
	// при передаче ptr - значение актуальное
	// при передаче значения - значение не обновляется
	i := 14
	defer _teardownValue(i) // закроется вторым
	defer _teardownPointer(&i) // закроется первым
	i++
	fmt.Println("Still working. i is", i)
	// Still working. i is 15
	// pointer closed, i is 15
	// value closed, i is 14

	var num int
	defer func(x int) {
		fmt.Println(x)
	}(num)
	num = 20
	// 0 (num уже вычислен для defer)
}

func _teardownValue (i int) {
	fmt.Println("value closed, i is", i)
}

func _teardownPointer (i *int) {
	fmt.Println("pointer closed, i is", *i)
}

