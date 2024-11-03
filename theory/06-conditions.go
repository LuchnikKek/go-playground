package theory

import (
	"fmt"
	"strconv"
)

func MainConditions () {
	mainConditionIf()
	mainConditionSwitch()
}

func mainConditionIf () {
	fmt.Println("============ if-elif-else ===========")
	// if-elif-else
	fmt.Println(_zeroChecker(-1), _zeroChecker(0), _zeroChecker(1))

	// Declaration in condition
	if x:= "kek"; x == "kek" {fmt.Println("yep, it's kek")}

	// Declaration in condition (multiple conditions checker)
	fmt.Println(_returningBlocks())
}

func _zeroChecker (i int) string {
	if i < 0 {
		return "less"
	} else if i > 0 {
		return "more"
	} else {
		return "equal"
	}
}

func _returningBlocks () string {
	const (
		MINIMAL_ONLINE = 100
		MAX_TIMEOUT = 10.0
	)
	
	_getCurrentOnline := func () int { return 1000 }
	_getCurrentTimeout := func () float64 { return 1.12 }

	if currentOnline := _getCurrentOnline(); currentOnline < MINIMAL_ONLINE {
		return "Слишком маленький онлайн: " + strconv.Itoa(currentOnline)
	} else if currentTimeout := _getCurrentTimeout(); currentTimeout > MAX_TIMEOUT {
		return "Слишком большой таймаут запросов: " + fmt.Sprintf("%f", currentTimeout)
	} else {
		return "Всё ок"
	}
}

func mainConditionSwitch () {
	fmt.Println("============ Switch-Case ===========")
	// switch по значению
	switchByValue(1)
	switchByValue(3)

	// switch по условию
	fmt.Println(switchByCondition(-1))
	fmt.Println(switchByCondition(0))
	fmt.Println(switchByCondition(1))
}

func switchByValue (value int) {
	switch value {
	case 1:
		value += 100
		fmt.Println("case 1:", value)
	default:
		fmt.Println("no such case")
	}
}

func switchByCondition (i int) (res string) {
	switch {
	case i < 0:
		res = "less"
	case i > 0:
		res = "more"
	case i == 0:
		res = "equal"
	}
	return
}
