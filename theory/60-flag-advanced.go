package theory

import (
	"flag"
	"fmt"
	"os"
)

func MainStartupFlagsAdvanced() {
	printArgs()
	_ = parseArgs // отрубил вызов, конфликтует с функцией ниже
	_ = parseArgsStruct
	parseArgsPositioned()
}

// простое получение списка всех (например, утилита получает только кучу файлов)
func printArgs() {
	fmt.Printf("Command: %v\n", os.Args[0])

	for i, v := range os.Args[1:] {
		fmt.Println(i+1, v)
	}
	fmt.Println("=== Finished flags output ===")
}

// эквивалентно -file a, -file=a, --file a, --file=a

// Для bool флагов:
//   - Если значение не передано -flag будет true (по умолчанию)
//   - Допустимые значения: 1, 0, t, f, true, false
//   - Нельзя -flag true, можно -flag=true
func parseArgs() {
	// go run main.go -file lol.png -dest ./output/result -w 256 -thumb

	// имя флага, значение по умолчанию, описание
	imgFile := flag.String("file", "", "input image file")
	destDir := flag.String("dest", "./output", "destination folder")
	width := flag.Int("w", 1024, "width of the image")
	isThumb := flag.Bool("thumb", false, "create thumb")

	// flag.String() возвращает указатель
	flag.Parse() // до flag.Parse() флагам присвоены значения по умолчанию

	fmt.Println("Image file:", *imgFile)
	fmt.Println("Destination folder:", *destDir)
	fmt.Println("Width:", *width)
	fmt.Println("Thumbs:", *isThumb)
}

// Те же яйца, только в структуру
func parseArgsStruct() {
	var options struct {
		file  string
		dest  string
		width int
		thumb bool
	}

	// связывать так
	flag.StringVar(&options.file, "file", "", "input image file")
	flag.StringVar(&options.dest, "dest", "./output", "destination folder")
	flag.IntVar(&options.width, "w", 1024, "width of the image")
	flag.BoolVar(&options.thumb, "thumb", false, "create thumb")
	// если передать два одинаковых указателя в разные флаги -
	// будет взято ПОСЛЕДНЕЕ из CLI

	flag.Parse()

	fmt.Println("Image file:", options.file)
	fmt.Println("Destination folder:", options.dest)
	fmt.Println("Width:", options.width)
	fmt.Println("Thumbs:", options.thumb)
}

// go run main.go -dest ./output/result -w 256 -thumb foo.png bar.png
func parseArgsPositioned() {
	destDir := flag.String("dest", "./output", "destination folder")
	width := flag.Int("w", 1024, "width of the image")
	isThumb := flag.Bool("thumb", false, "create thumb")

	// flag.String() возвращает указатель
	flag.Parse() // до flag.Parse() флагам присвоены значения по умолчанию

	// Вся позиционка летит в flag.Args(). flag.Arg(i) - вернёт i-ый
	for i, v := range flag.Args() {
		fmt.Printf("Image file (%d): %s\r\n", i, v)
	}

	fmt.Println("Destination folder:", *destDir)
	fmt.Println("Width:", *width)
	fmt.Println("Thumbs:", *isThumb)
}
