package theory

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func MainFilesRead() {
	readFull()    // файл целиком в оперативу
	readByChunk() // файл по частям
}

const readFilename = "theory/35-read-example.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFull() {
	data, err := os.ReadFile(readFilename)
	check(err)
	fmt.Printf("full length: %v\n", len(string(data)))
}

func readByChunk() {
	f, err := os.Open(readFilename)
	check(err)

	// Read - читает первые N байт
	b1 := make([]byte, 7)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1)) // 7 bytes: abcdefg

	// Read - затем следующие
	b2 := make([]byte, 3)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes: %v\n", n2, string(b2)) // 3 bytes: hij

	// Seek()
	o3, err := f.Seek(1, 0) // 3 bytes @ 1: bcd  ---- скипнет 1, от начала
	// o3, err := f.Seek(1, 1) // 3 bytes @ 11: lmn ---- скипнет 1 от текущего, будет lmn
	// o3, err := f.Seek(-3, 2) // 3 bytes @ 70: one ---- скипнет три с конца? их же считает (порядок верный)
	check(err)
	b3 := make([]byte, 3)
	n3, err := io.ReadAtLeast(f, b3, 2) // ReadAtLeast - читает первые N байт, но не меньше 2
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3)) // bcd

	// Peek() - читает N байтов, но не двигает курсор дальше
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(13)
	check(err)
	fmt.Printf("13 bytes: %s\n", string(b4))

	f.Close()
}
