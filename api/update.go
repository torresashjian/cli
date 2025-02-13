package api

import (
	"fmt"
	"os/exec"

	"github.com/torresashjian/cli/common"
	"github.com/torresashjian/cli/util"
)

func UpdatePkg(project common.AppProject, pkg string) error {

	if Verbose() {
		fmt.Printf("Updating Package: %s \n", pkg)
	}

	err := util.ExecCmd(exec.Command("go", "get", "-u", pkg), project.SrcDir())
	return err
}
