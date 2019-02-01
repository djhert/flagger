package flagger

import (
	"fmt"
)

const (
	NAME     = "flagger"
	MAJORVER = "0"
	MINORVER = "8"
	VERTAG   = "-alpha"
	DESC     = "Flag handler"
)

// Version returns a formatted string of the name/version number
func Version() string {
	return fmt.Sprintf("%s v%s.%s%s", NAME, MAJORVER, MINORVER, VERTAG)
}

// Info returns a formatted string of Version and the Description
func Info() string {
	return fmt.Sprintf("%s\n\t%s", Version(), DESC)
}
