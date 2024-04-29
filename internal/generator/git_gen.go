package generator

import (
	"os"
	"os/exec"
)

const gitignoreTemplate = `
*.exe
*.exe~
*.dll
*.so
*.dylib

*.test

*.out

vendor/

go.work
`

func createGitLocalRepo() (err error) {
	if err = os.WriteFile(
		".gitignore",
		gitignoreFileText(),
		os.ModePerm); err != nil {
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		return
	}

	if err = initRepo(pwd); err != nil {
		return
	}

	if err = createInitCommit(pwd); err != nil {
		return
	}

	return
}

func gitignoreFileText() []byte {
	return []byte(gitignoreTemplate)
}

func initRepo(pwd string) (err error) {
	cmd := exec.Command("git", "init")
	cmd.Dir = pwd
	return cmd.Run()
}

func createInitCommit(pwd string) (err error) {
	cmdAdd := exec.Command("git", "add", ".")
	cmdAdd.Dir = pwd
	if err = cmdAdd.Run(); err != nil {
		return
	}

	cmdCommit := exec.Command("git", "commit", "-m", "Init")
	cmdCommit.Dir = pwd
	if err = cmdCommit.Run(); err != nil {
		return
	}

	return
}
