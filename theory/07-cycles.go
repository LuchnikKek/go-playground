package main

import "fmt"

func main () {
	mainInfiniteLoop()
	mainConditionLoop()
	mainParametrizedLoop()
	mainRangedLoop()
	mainStrIterating()
}

const ITERS_COUNT = 3

func mainInfiniteLoop () {
	fmt.Println("===== InfiniteLoop =======")
	i := 0
	for {
		fmt.Println("Итерация:", i)
		i++
		if i == ITERS_COUNT {break}
	}
}

func mainConditionLoop () {
	fmt.Println("===== ConditionLoop =======")
	i := 0
	for i < ITERS_COUNT {
		fmt.Println("Итерация:", i)
		i++
	}
}

func mainParametrizedLoop () {
	fmt.Println("===== ParametrizedLoop =======")
	for i := 0; i < ITERS_COUNT; i++ {
		fmt.Println("Итерация:", i)
	}
}

func mainRangedLoop () {
	fmt.Println("===== RangedLoop =======")
	for i := range ITERS_COUNT {
		fmt.Println("Итерация:", i)
	}

	// range by slice
	sl := []int{5,10,15}
	for i, value := range sl {
		fmt.Printf("In slice: key %v value %v\n", i, value)
	}

	// range by map
	dict := map[int]int {
		5: 100,
		6: 600,
		7: 52,
	}
	for key, value := range dict {
		fmt.Printf("In map: key %v value %v\n", key, value)
	}
}

func mainStrIterating() {
	// Два способа итерирования по строке (байты и символы)
	fmt.Println("===== Strings Iterating =======")
	str := "Кот"

	// parametrized by string (index 0,1,2,3,4,5; value type uint8=byte)
	fmt.Println("By parametrized (by bytes):")
	for i := 0; i < len(str); i++ {
		fmt.Printf("index=%v, value=%v, type=%T\n", i, str[i], str[i])
	}

	// range by string (index 0,2,4; values type int32=rune)
	fmt.Println("By range (by symbols):")
	for idx, value := range str {
		fmt.Printf("index=%v, value=%v, type=%T\n", idx, value, value)
	}
}