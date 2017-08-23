package builder

import (
	"github.com/michaelwmerritt/project-builder/model"
	"os"
	"os/exec"
	"fmt"
	"bufio"
)

type SvnManager struct {}

func NewSvnManager() SvnManager {
	return SvnManager{}
}

func (svnManager SvnManager) Checkout(release model.Release, modules []model.Module, path string) (err error) {
	svnUsername := ""
	svnPassword := ""
	svnCommand := "svn"
	svnCheckoutCommand := "co"

	for _, module := range modules {
		cmd := exec.Command(svnCommand, svnCheckoutCommand, "--username", svnUsername, "--password", svnPassword, module.CreateVCSUrl(), module.CreateModulePath(path))
		err = executeCommand(cmd)
	}
	return
}

func executeCommand(cmd *exec.Cmd) (err error) {
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()
	err = cmd.Run()
	return
}
