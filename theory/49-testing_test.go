// Команды:
//   - go test -v . -count 1								запуск без кеша
//   - go clean -testcache									Очистка кеша
//   - go test -run Sum										запуск, где в имени есть Sum
//   - go test -run Sum$									запуск, где имя кончается на Sum
//   - go test -run ^TestSum$								запуск с конкретным именем

// Подтесты:
//   - go test -v -run "Sum/^with negative values$"			запуск с именем with negative values
//   - go test -v -run "Sum/negative"						запуск с negative в имени
package theory

import "testing"

func TestSum(t *testing.T) {
	if sum := Sum(1, 2); sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestSumParametrized(t *testing.T) {
	tests := []struct { // добавляем слайс тестов
		name   string
		values []int
		want   int
	}{
		{
			name:   "simple test #1", // описываем каждый тест:
			values: []int{1, 2},      // значения, которые будет принимать функция,
			want:   3,                // и ожидаемый результат
		},
		{
			name:   "one",
			values: []int{1},
			want:   1,
		},
		{
			name:   "with negative values",
			values: []int{-1, -2, 3},
			want:   0,
		},
		{
			name:   "with negative zero",
			values: []int{-0, 3},
			want:   3,
		},
		{
			name: "a lot of values",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
				14, 15, 16, 17, 18, 18},
			want: 189,
		},
	}
	for _, test := range tests { // цикл по всем тестам
		t.Run(test.name, func(t *testing.T) {
			if sum := Sum(test.values...); sum != test.want {
				t.Errorf("Sum() = %d, want %d", sum, test.want)
			}
		})
	}
	// for _, test := range tests { // то же, но без запуска в отдельном тесте
	// 	if sum := Sum(test.values...); sum != test.want {
	// 		t.Errorf("%s: Sum() = %d, want %d", test.name, sum, test.want)
	// 	}
	// }
}
