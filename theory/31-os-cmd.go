package theory

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func MainOsExec() {
	execCmd()
	execWithArgs()
	execWithVarArgs()
}

func execCmd() {
	cmd := exec.Command("ls")

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	// Тут будет ждать, т.к. в cmd.Run() встроен .wait(),
	// который ждёт передачи дескриптора от приложения
}

func execWithArgs() {
	cmd := exec.Command("tr", "a-z", "A-Z")

	cmd.Stdin = strings.NewReader("shawty pass out")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("result phrase: %q\n", out.String())
	// result phrase: "SHAWTY PASS OUT"
}

func execWithVarArgs() {
	arg1 := "foo"
	arg2 := "bar"
	arg3 := "buz"

	// stdout, err := exec.Command("echo", arg1, arg2, arg3).Output()
	cmd := exec.Command("echo", arg1, arg2, arg3)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	log.Printf("result phrase: %q\n", string(stdout))
	// result phrase: "foo bar buz\n"
}
