package commands

import (
	"errors"
	"fmt"
	"github.com/hlfstr/flagger"
)

var (
	NoCmds = errors.New("No Commands passed")
)

func New() *Commands {
	c := &Commands{}
	c.cmds = make(map[string]Commander)
	return c
}

type Commander interface {
	Prepare(*flagger.Flags)
	Action([]string, *flagger.Flags) error
	Print()
}

type Commands struct {
	cmds map[string]Commander
}

func (c *Commands) Add(cmd string, cmdr Commander) {
	c.cmds[cmd] = cmdr
}

func (c *Commands) Parse(flags []string) error {
	if len(flags) > 0 {
		if v, ok := c.cmds[flags[0]]; ok {
			f := flagger.New()
			v.Prepare(f)
			return v.Action(flags[1:], f)
		}
		return fmt.Errorf("Unable to locate command %s", flags[0])
	}
	return NoCmds
}

func (c *Commands) Print() {
	fmt.Println("Available Commands:")
	for i := range c.cmds {
		c.cmds[i].Print()
	}
}
