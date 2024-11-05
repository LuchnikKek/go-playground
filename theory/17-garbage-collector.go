package theory

import "runtime"

type Photo struct {
	height int
	width int
	count int
}

func NewPhoto(height int, width int) Photo {
	// Возвращает копию структуры. Всё, что было в памяти стека будет удалено
	return Photo{height: height, width: width, count: 1}
}

func NewPhotoRef(height int, width int) *Photo {
	// Резервирует память для структуры. Возвращает ссылку на структуру.
	// Это будет в куче. Память резервируется в куче и не освобождается.
	// Вернее, освобождается, но Garbage Collector'ом.
	return &Photo{height: height, width: width, count: 1}
	// GC запускается конкуррентно, в фоне.
	// Чем больше используется аллокация в Heap (куче), тем чаще будет запускаться GC
	// Stop The World - пока сборщик анализирует, приложение ожидает
	// Использует трехцветный алгоритм пометки и очистки для выбора. Потом удаляет
}

func MainGarbageCollector() {
	runtime.GC()
	// запуск GC
}
