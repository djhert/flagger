package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/hlfstr/flagger"
	"github.com/hlfstr/flagger/commands"
)

type cmdA struct {
	b bool
	i int
	s string
}

func (c *cmdA) Prepare(flags *flagger.Flags) {
	flags.AddHelp("CommandA Help", "Show CommandA flags")
	flags.BoolVar(&c.b, "Test Bool", "-b", "--bool")
	flags.IntVar(&c.i, 4, "Test Int", "-i", "--integer")
	flags.StringVar(&c.s, "cmdA", "Test cmdA String", "-s", "--string")

}

func (c *cmdA) Action(s []string, f *flagger.Flags) error {
	fmt.Println("cmdA Action")
	fmt.Printf("Before\n")
	fmt.Printf("  Bool:   %t\n", c.b)
	fmt.Printf("  Int:    %d\n", c.i)
	fmt.Printf("  String: %s\n", c.s)
	data, err := f.Parse(s)
	if err != nil {
		return err
	}
	fmt.Println("After")
	fmt.Printf("  Bool:   %t\n", c.b)
	fmt.Printf("  Int:    %d\n", c.i)
	fmt.Printf("  String: %s\n", c.s)
	fmt.Print("Data: ")
	fmt.Println(data)
	return nil
}

type cmdB struct {
	b bool
	i int
	s string
}

func (c *cmdB) Prepare(flags *flagger.Flags) {
	flags.AddHelp("CommandB Help", "Show CommandB flags")
	flags.BoolVar(&c.b, "Test Bool", "-b", "--bool")
	flags.IntVar(&c.i, 9, "Test Int", "-i", "--integer")
	flags.StringVar(&c.s, "cmdB", "Test cmdB String", "-s", "--string")
}

func (c *cmdB) Action(s []string, f *flagger.Flags) error {
	fmt.Println("cmdB Action")
	fmt.Printf("Before\n")
	fmt.Printf("  Bool:   %t\n", c.b)
	fmt.Printf("  Int:    %d\n", c.i)
	fmt.Printf("  String: %s\n", c.s)
	data, err := f.Parse(s)
	if err != nil {
		return err
	}
	fmt.Println("After")
	fmt.Printf("  Bool:   %t\n", c.b)
	fmt.Printf("  Int:    %d\n", c.i)
	fmt.Printf("  String: %s\n", c.s)
	fmt.Print("Data: ")
	fmt.Println(data)
	return nil
}

func main() {
	commands.New()
	commands.AddHelp("Available Commands:\n")
	commands.AddVersion(flagger.Version())
	a := cmdA{}
	b := cmdB{}
	commands.Add("one", &a)
	commands.Add("two", &b)
	err := commands.Parse(os.Args)
	if errors.Is(err, commands.ErrNoCmds) {
		commands.Usage("[COMMAND] [OPTION]...", flagger.Info())
	} else if errors.Is(err, commands.ErrBadCmd) {
		commands.Usage("[COMMAND] [OPTION]...", err.Error())
		os.Exit(1)
	} else if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
