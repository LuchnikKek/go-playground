package main

import "fmt"

func main() {
	mainArrays()
	mainSlices()
	mainMaps()
	mainStructs()
}

func mainArrays() {
	// ========= Массивы ===========
	fmt.Println("========= Массивы ===========")
	// Имеют фиксированную длину
	var arr1 [5]int
	fmt.Println(arr1) // [0 0 0 0 0]

	// объявление
	var arr2 [3]int = [3]int{1,2,3} // излишняя форма
	var arr3 = [3]int{1,2,3} // полная форма
	arr4 := [3]int{1,2,3} // краткая форма
	fmt.Println(arr2, arr3, arr4)

	// размер массива исходя из кол-ва элементов в объявлении
	arr5 := [...]int{1,2}
	fmt.Println("Элементов в массиве:", len(arr5)) // 2
}

func mainSlices() {
	// ========== Слайсы, массивы динамического размера ============
	fmt.Println("========= Слайсы, срезы =========")

	// Имеют переменную длину, элементы одного типа
	someSlice := []int{10,20,30,40,50}
	newSlice := someSlice[2:4]
	fmt.Println(newSlice) // 30, 40
	fmt.Println(len(newSlice)) // 2 - количество элементов
	fmt.Println(cap(newSlice)) // 3 - capacity, ёмкость, кол-во под которое выделена память

	// Создание среза - make 
	newSlice = make([]int, 1, 2)
	_printSliceInfo(newSlice) // len=1, cap=2, arr=[0]

	newSlice = append(newSlice, 2)
	_printSliceInfo(newSlice) // len=2, cap=2, arr=[0 2]

	// как только len > cap, создаётся новый массив (cap * 2), значения копируются
	newSlice = append(newSlice, 3)
	_printSliceInfo(newSlice) // len=3, cap=4, arr=[0 2 3]
}

func _printSliceInfo(sl []int) {
	fmt.Printf("len=%d, cap=%d, arr=%v\n", len(sl), cap(sl), sl)
}

func mainMaps() {
	// ========= Maps (хэш-таблицы) ============
	// Все ключи должны быть сравниваемыми, без коллизий (float не подойдёт)
	fmt.Println("========= Maps =========")

	// Инициализация
	m1 := make(map[int]int) // map[]
	_ = m1

	// map[тип ключа]тип значения
	var m2 map[int]bool
	fmt.Println(m2)

	ages := map[string]uint16 {
		"Гена": 30,
		"Вася": 1,
	}
	fmt.Println(ages) // map[Вася:1 Гена:30 Илья:52]

	ages["Илья"] = 52 // create

	fmt.Println(ages["Илья"]) // read

	ages["Илья"]++ // update
	fmt.Println(ages["Илья"])
}

func mainStructs() {
	// ============ Структуры ============
	fmt.Println("========= Структуры =========")

	// объявление структуры
	type Point struct {
		x int
		y int
	}

	p1 := Point{x: 1, y: 2} // {1 2}
	fmt.Println(p1)

	p2 := Point{3, 4} // {3, 4}
	p2 = Point{3, 5} // {3, 5}
	p2.x = 30
	fmt.Println(p2) // {30, 5}
}
