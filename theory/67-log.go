package theory

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func logInFileExample() {
	// Создаем файл, обрабатываем ошибку
	file, err := os.OpenFile("./theory/67-info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// отложенное закрытие файла
	defer file.Close()

	// устанавливаем назначение вывода логов в файл
	//log.SetOutput(file)
	//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//log.Println("Hello, World in file!")
	//2025/06/22 18:07:51 Hello, World in file!

	// то же самое
	// LstdFlags (дата, время), Lshortfile (имя файла, строка)
	mylog := log.New(file, `serv `, log.LstdFlags|log.Lshortfile)
	mylog.Println("Hello, World in file!")
	//serv 2025/06/22 18:11:53 67-log.go:24: Hello, World in file!
}

func logInBufferExample() {
	var buf bytes.Buffer

	bufLog := log.New(&buf, "bufLog: ", 0)
	bufLog.Println("Hello, World in buffer!")
	bufLog.Println("Bye!")

	fmt.Println(&buf)
}

func MainLogsAdvanced() {
	logInFileExample()
	logInBufferExample()
}
