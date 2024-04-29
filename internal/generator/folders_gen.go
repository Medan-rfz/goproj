package generator

import (
	"fmt"
	"os"

	"goproj/internal/configs"
)

const (
	internalFolder = "internal"
	pkgFolder      = "pkg"
)

func CreateStructure(config *configs.StructureConfig) (err error) {
	cmdAppFolder := fmt.Sprintf("cmd/%s", config.AppName)

	if err = os.MkdirAll(cmdAppFolder, os.ModePerm); err != nil {
		return
	}

	if err = os.MkdirAll(pkgFolder, os.ModePerm); err != nil {
		return
	}

	if err = os.MkdirAll(internalFolder, os.ModePerm); err != nil {
		return
	}

	if err = os.WriteFile(
		fmt.Sprintf("%s/main.go", cmdAppFolder),
		mainFileText(),
		os.ModePerm); err != nil {
		return
	}

	if err = os.WriteFile(
		"Makefile",
		makefileText(config),
		os.ModePerm); err != nil {
		return
	}

	if config.GitExpect {
		if err = createGitLocalRepo(); err != nil {
			return
		}
	}

	return
}
