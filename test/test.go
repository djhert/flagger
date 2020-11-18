package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/hlfstr/flagger"
)

func main() {
	f := flagger.New()
	f.AddHelp("Show this help", "Testing Flags:")
	f.AddVersion("Show version", flagger.Info())

	//Bool flag
	b := f.Bool("Test Bool", "-b", "--bool")

	//int flag
	i := f.Int(8, "Test Int", "-i", "--integer")

	//string flag
	s := f.String("h", "Test String", "-s", "--string")

	fmt.Printf("Before\n")
	fmt.Printf("  Bool:   %t\n", *b)
	fmt.Printf("  Int:    %d\n", *i)
	fmt.Printf("  String: %s\n", *s)
	d, err := f.Parse(os.Args)
	if errors.Is(err, flagger.ErrNoFlags) {
		f.Usage("[OPTION]...", fmt.Sprintf("Try '%s --help' for more information", f.Name))
		os.Exit(0) // Exit clean
	} else if err != nil {
		f.Usage("[OPTION]...", err.Error())
		fmt.Printf("Try '%s --help' for more information\n", f.Name)
		os.Exit(1) // Exit error
	}
	fmt.Println("After")
	fmt.Printf("  Bool:   %t\n", *b)
	fmt.Printf("  Int:    %d\n", *i)
	fmt.Printf("  String: %s\n", *s)
	fmt.Print("Data: ")
	fmt.Println(d)
}
