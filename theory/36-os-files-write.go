package theory

import (
	"bufio"
	"fmt"
	"os"
)

const (
	writeFilename = "theory/36-write-example.txt"
	writeMessage  = "Smells Like Teen Spirit"
)

func MainFilesWrite() {
	// writeFull()
	writeExamples()
}

// func writeFull() {
// 	// WriteFile записывает, но зачищает файл
// 	d1 := []byte(writeMessage)
// 	err := os.WriteFile(writeFilename, d1, 0644)
// 	check(err)
// }

func writeExamples() {
	// os.Create() - Создание файла
	f, err := os.Create(writeFilename)
	check(err)

	defer f.Close()

	// Write() - пишет байтики
	d2 := []byte{83, 109, 101, 108, 108, 115, 32, 76, 105, 107, 101, 32, 84, 101, 101, 110, 32, 83, 112, 105, 114, 105, 116, 10} // 10 = \n
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString() - пишет string, просто обёртка над Write()
	n3, err := f.WriteString("Overbored\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)
	// .Sync() гарантирует, что файл записался на диск
	_ = f.Sync()

	// Поток (Writer), создаётся от файла и мы можем туда чет записать
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("Self assured\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush() // Запишет всё в файл
}
