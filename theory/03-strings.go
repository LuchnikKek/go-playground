package theory

import "fmt"

func mainStrings() {
	// Immutables!

	var (
		renderedStr = "wow\tcool\n" // ""
		nonRenderedStr = `wow\tcool\n` // ``
	)
	fmt.Println("Renders:", renderedStr)
	fmt.Println("Not renders:", nonRenderedStr)

	// ''
	var b byte = 'r' // 1 byte ASCII
	var r rune = '你' // 3 byte UTF
	fmt.Println(b, r)

	twoBytesStr := "ю" // кодируется двумя байтами
	fmt.Println(twoBytesStr[0]) // 209
	fmt.Println(twoBytesStr[1]) // 142
	fmt.Println(len(twoBytesStr)) // 2

	var someStr string = "абвгд"
	fmt.Println(len(someStr)) // 10
	fmt.Println(someStr[:4]) // аб
	fmt.Println(someStr[4:]) // вгд
	
	// concat
	fmt.Println(someStr[:3] + someStr[3:]) // а? + ?вгд = абвгд
	
	// сравнение
	fmt.Println("aб" < "ав") // true, лексикографический, как в словаре

	someText := "какой-то текст."
	some := someText[:15]
	text := someText[16:26]
	fmt.Println(some)
	fmt.Println(text)
	// len(someText) - количество бит. someText[len(someText)] -> panic
}
