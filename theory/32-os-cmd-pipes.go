package theory

import (
	"io"
	"log"
	"os/exec"
)

func MainOsExecPipes() {
	inputPipe()
	outputPipe()
}

func inputPipe() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		_, _ = io.WriteString(stdin, "babushka boy")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%q\n", string(out))
}

func outputPipe() {
	// Удобно, когда нужно запустить логгер где-нибудь в коро/главном
	// Чтобы читать из него
	cmd := exec.Command("echo", "a nichego tot fact chto")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	log.Printf("%q\n", string(data))
}
