package cmds

import (
	"io/ioutil"
	"strings"

	"github.com/abiosoft/ishell"
)

func View(c *ishell.Context) {
	file := strings.Join(c.Args, " ")
	output, err := ioutil.ReadFile(file)
	if err != nil {
		c.Println(err.Error())
		return
	}
	c.ShowPaged(string(output))
}
