package commands

import (
	"errors"
	"fmt"

	"github.com/hlfstr/flagger"
)

var (
	ErrNoCmds  = errors.New("No Commands passed")
	ErrNoInit  = errors.New("Commands not initialized")
	ErrBadCmd  = errors.New("invalid command")
	ErrHelp    = errors.New("Help")
	ErrVersion = errors.New("Version")
)

var com *Commands

// Commander interface defines how a Command should operate
type Commander interface {
	Prepare(*flagger.Flags)
	Action([]string, *flagger.Flags) error
}

type Commands struct {
	cmds map[string]Commander
	Name string
}

func ifInit() error {
	if com == nil {
		return ErrNoInit
	}
	return nil
}

func New() {
	com = &Commands{}
	com.cmds = make(map[string]Commander)
	com.Name = ""
}

func Name() string {
	if err := ifInit(); err != nil {
		return ""
	}
	return com.Name
}

func Add(cmd string, cmdr Commander) error {
	if err := ifInit(); err != nil {
		return err
	}
	com.cmds[cmd] = cmdr
	return nil
}

func Parse(flags []string) error {
	if err := ifInit(); err != nil {
		return err
	}
	com.Name = flags[0]
	if len(flags) > 1 {
		if v, ok := com.cmds[flags[1]]; ok {
			f := flagger.New()
			v.Prepare(f)
			return v.Action(flags[2:], f)
		}

		return fmt.Errorf("%s: %w -- '%s'", com.Name, ErrBadCmd, flags[1])
	}
	return ErrNoCmds
}

func Usage(usage string, msg string) error {
	if err := ifInit(); err != nil {
		return err
	}
	fmt.Printf("Usage: %s %s\n", com.Name, usage)
	fmt.Printf(msg)
	fmt.Printf("\nAvailable Commands:\n")
	print(false)
	return nil
}

func print(full bool) {
	for i := range com.cmds {
		if full {
			f := flagger.New()
			com.cmds[i].Prepare(f)
			f.Help(fmt.Sprintf(" %s", i))
			fmt.Println()
		} else {
			fmt.Printf(" %s\n", i)
		}
	}
}

func Help(msg string) error {
	if err := ifInit(); err != nil {
		return err
	}
	fmt.Printf("Usage: %s [COMMAND] [OPTION]...\n", com.Name)
	fmt.Printf(msg)
	print(true)
	return ErrHelp
}

type help struct {
	txt string
}

func (h *help) Prepare(flags *flagger.Flags) {}

func (h *help) Action(s []string, f *flagger.Flags) error {
	return Help(h.txt)
}

func AddHelp(msg string) error {
	if err := ifInit(); err != nil {
		return err
	}
	h := &help{txt: msg}
	Add("help", h)
	return nil
}

type version struct {
	txt string
}

func (v *version) Prepare(flags *flagger.Flags) {}

func (v *version) Action(s []string, f *flagger.Flags) error {
	fmt.Println(v.txt)
	return ErrVersion
}

func AddVersion(msg string) error {
	if err := ifInit(); err != nil {
		return err
	}
	v := &version{txt: msg}
	Add("version", v)
	return nil
}
