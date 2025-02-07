package theory

import (
	"bytes"
	"encoding/gob"
	"log"
)

func MainGob() {
	decodeGob(encodeGob())
}

func encodeGob() []byte {
	var buf bytes.Buffer        // создаем буфер
	enc := gob.NewEncoder(&buf) // создаем байтовый энкодер (кодирует представление в битовый формат)

	m := make(map[string]string) // структура для примера (название не сохранится при передаче)
	m["foo"] = "bar"             // сохранятся ток поля и типы этих полей

	if err := enc.Encode(m); err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}

func decodeGob(input []byte) {
	buf := bytes.NewBuffer(input)
	dec := gob.NewDecoder(buf)

	m := make(map[string]string) // тот же тип, что и отправляемая структура

	if err := dec.Decode(&m); err != nil {
		log.Fatal(err)
	}

	log.Printf("value %s\n", m["foo"])
}
