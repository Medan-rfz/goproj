package generator

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
)

func CreateBaseProject(name string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	errBuf := bytes.Buffer{}
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = pwd
	cmd.Stderr = &errBuf

	if err := cmd.Run(); err != nil {
		return errors.New(errBuf.String())
	}

	return nil
}
