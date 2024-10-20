package theory

import "fmt"

func mainTypes() {
	// basic types: integer, float, complex, bool, string
	// aggregate types: array, slice, map, struct

	// ======== Числовой тип ==========
	// целые числа
	const signedInt8Bot int8 = -128
	const signedInt8Top int8 = 127 // math.Pow(2, 8 - 1) - 1
	fmt.Println("Значения в диапазоне [-128, 127]")

	// целые беззнаковые числа
	const unsignedInt8Bot uint8 = 0
	const unsignedInt8Top uint8 = 255 // math.Pow(2, 8) - 1
	fmt.Println("Значения в диапазоне [0, 255]", unsignedInt8Top)

	// переполнение uint
	var overflowedUint uint = 1 // uint32/uint64 (в зависимости от арх-ры процессора)
	overflowedUint -= 2
	fmt.Println(overflowedUint) // 18446744073709551615

	// byte == uint8 - байты
	// rune == int32 - символы юникод
	// см. 03-strings.go

	// с плавающей точкой ~8 знаков
	fmt.Println(float32(1) / 7) // 0.14285715
	
	// с плавающей точкой ~17 знаков
	fmt.Println(float64(1) / 7) // 0.14285714285714285

	var buggedFloat32 float32 = 16_777_216 // опасное float32
	fmt.Println(buggedFloat32 == buggedFloat32 + 1) // true

	// комплексные числа
	var compl1 complex64;
	var compl2 complex128;
	_, _ = compl1, compl2


	// ======== Логический тип ==========
	trueNotFalse := !false
	trueAnd := true && 10 > 1 
	trueOr := true || false 
	fmt.Println(trueNotFalse, trueAnd, trueOr)

	fmt.Println("0 is", _intToBool(0))
	fmt.Println("5 is", _intToBool(5))
}

func _intToBool(i int) bool {
	if i != 0 {
		return true
	}
	return false
}
