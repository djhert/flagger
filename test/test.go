package main

import (
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
	d, err := f.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Try '%s --help' for more information\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println("After")
	fmt.Printf("  Bool:   %t\n", *b)
	fmt.Printf("  Int:    %d\n", *i)
	fmt.Printf("  String: %s\n", *s)
	fmt.Print("Data: ")
	fmt.Println(d)
}
