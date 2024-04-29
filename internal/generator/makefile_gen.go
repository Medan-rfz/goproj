package generator

import (
	"fmt"

	"goproj/internal/configs"
)

const makefileTemplate = `
build:
	go build -o ./bin/ ./cmd/%s

run:
	go run ./cmd/%s
`

func makefileText(config *configs.StructureConfig) []byte {
	return []byte(fmt.Sprintf(makefileTemplate, config.AppName, config.AppName))
}
