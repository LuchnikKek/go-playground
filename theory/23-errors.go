package theory

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ServerError struct {
	err error
}

func (s *ServerError) Error() string {
	return s.err.Error()
}

func NewServerError(msg string) error {
	return &ServerError{err: errors.New(msg)}
}

func getClient() (string, error) {
	err := NewServerError("2 - такая вот ошибка")
	return "client", err
}

func getJson(msg []byte) (int, error) {
	var n int
	err := json.Unmarshal(msg, &n)
	if err != nil {
		return 0, fmt.Errorf("ошибка распаковки json: %w", DecoderError) // wraps
	}
	return n, nil
}

var DecoderError = NewServerError("ошибка декодера")

func MainErrors() {
	err := errors.New("1 - такая вот ошибка")
	fmt.Println(err.Error())

	cli, err := getClient()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(cli)
	}

	n, err := getJson([]byte(`1asd1`))
	if errors.Is(err, DecoderError) {
		fmt.Println(err.Error())
		// nextErr := errors.Unwrap(err)
		// fmt.Println("wrapped:", nextErr)
	} else {
		fmt.Printf("msg is %v\n", n)
	}
}
