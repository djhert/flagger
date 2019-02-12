package main

import (
	"fmt"
	"github.com/hlfstr/flagger"
	"github.com/hlfstr/flagger/commands"
	"os"
)

type cmdA struct {
	b bool
	i int
	s string
}

func (c *cmdA) Prepare(flags *flagger.Flags) {
	flags.BoolVar(&c.b, "Test Bool", "-b", "--bool")
	flags.IntVar(&c.i, 4, "Test Int", "-i", "--integer")
	flags.StringVar(&c.s, "cmdA", "-s", "--string")
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

func (c *cmdA) Print() {
	fmt.Printf("  CMDA \t Command A Help goes here.\n")
}

type cmdB struct {
	b bool
	i int
	s string
}

func (c *cmdB) Prepare(flags *flagger.Flags) {
	flags.BoolVar(&c.b, "Test Bool", "-b", "--bool")
	flags.IntVar(&c.i, 9, "Test Int", "-i", "--integer")
	flags.StringVar(&c.s, "cmdB", "-s", "--string")
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

func (c *cmdB) Print() {
	fmt.Printf("  CMDB \t Command B Help goes here.\n")
}

func main() {
	cmd := commands.New()
	a := cmdA{}
	b := cmdB{}
	cmd.Add("one", &a)
	cmd.Add("two", &b)
	err := cmd.Parse(os.Args[1:])
	if err == commands.NoCmds {
		cmd.Print()
	} else if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
