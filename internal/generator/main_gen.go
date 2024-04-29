package generator

const mainTemplate = `package main

func main() {

}
`

func mainFileText() []byte {
	return []byte(mainTemplate)
}
