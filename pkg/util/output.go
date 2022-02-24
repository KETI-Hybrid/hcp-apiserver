package util

import (
	"encoding/json"
	"io"
	"log"
	"os/exec"
)

func CombinedOutput2(cmd *exec.Cmd) ([]byte, []byte) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	errb, _ := io.ReadAll(stderr)
	outb, _ := io.ReadAll(stdout)
	cmd.Wait()
	return errb, outb
}

func GetOutput(cmd *exec.Cmd) ([]byte, error) {
	errb, outb := CombinedOutput2(cmd)
	output := Output{
		Stderr: errb,
		Stdout: outb,
	}
	data, err := json.Marshal(output)
	return data, err
}
