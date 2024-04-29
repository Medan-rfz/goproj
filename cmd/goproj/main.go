package main

import (
	"flag"

	"goproj/internal/configs"
	"goproj/internal/generator"

	"github.com/sirupsen/logrus"
)

func main() {
	appConfig := configs.GetStructConfigFromFlags()

	logrus.SetLevel(logrus.ErrorLevel)
	if appConfig.IsDebug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	projNameArg := flag.Args()
	if len(projNameArg) != 1 {
		logrus.Fatalln("Incorrect project name")
	}

	appConfig.AppName = projNameArg[0]

	if err := generator.CreateBaseProject(appConfig.AppName); err != nil {
		logrus.Fatalf("Go mod init error: %v", err)
	}

	config := &configs.StructureConfig{
		AppName:   appConfig.AppName,
		GitExpect: appConfig.GitExpect,
	}
	generator.CreateStructure(config)
}
