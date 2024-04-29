package configs

import "flag"

func GetStructConfigFromFlags() *AppConfig {
	debugFlag := flag.Bool("d", false, "[Optional] Output debug log")
	gitFlag := flag.Bool("g", false, "[Optional] Init git local repository")

	flag.Parse()

	return &AppConfig{
		IsDebug:   *debugFlag,
		GitExpect: *gitFlag,
	}
}
