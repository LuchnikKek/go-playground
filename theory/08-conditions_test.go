package theory

import "testing"

// go test
// go test -v - подробно
// go test -cover - покрытие

// go test -coverprofile=cover.out - сгенерировать отчёт по coverage
// go tool cover -html=cover.out -o cover.html - преобразовать в html
// open cover.html - открыть html

func TestSwitchByConditionLess(t *testing.T) {
	res := switchByCondition(-10)
	if res != "less" {
		t.Fatal("must be \"less\"")
	}
}

func TestSwitchByConditionEqual(t *testing.T) {
	res := switchByCondition(0)
	if res != "equal" {
		t.Fatal("must be \"equal\"")
	}
}

func TestSwitchByConditionMore(t *testing.T) {
	res := switchByCondition(10)
	if res != "more" {
		t.Fatal("must be \"more\"")
	}
}
