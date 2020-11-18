package flagger

import (
	"fmt"
)

const (
	// NAME of flagger
	NAME = "flagger"
	// MAJORVERSION of flagger
	MAJORVERSION = "1"
	// MINORVERSION of flagger
	MINORVERSION = "1"
	// VERSIONTAG of flagger
	VERSIONTAG = "0"
	// DESCRIPTION of flagger
	DESCRIPTION = "Flag handler written in Go"
)

// Version returns a formatted string of the name/version number
func Version() string {
	return fmt.Sprintf("v%s.%s.%s", MAJORVERSION, MINORVERSION, VERSIONTAG)
}

// Info returns a formatted string of Version and the Description
func Info() string {
	return fmt.Sprintf("%s %s\n -- %s", NAME, Version(), DESCRIPTION)
}
