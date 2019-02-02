package flagger

import (
	"fmt"
	"strconv"
)

type Flag struct {
	Usage string
	flags []string
	Value FlagValue
}

func (f Flag) string() string {
	s := fmt.Sprintf("  %s", f.flags[0])
	for i := 1; i < len(f.flags); i++ {
		s = fmt.Sprintf("%s, %s", s, f.flags[i])
	}
	return s
}

func (f Flag) Print() string {
	return fmt.Sprintf("%-25s\t%-s", f.string(), f.Usage)
}

type FlagValue interface {
	String() string
	Set(string) error
}

type Getter interface {
	FlagValue
	Get() interface{}
}

type boolFlag bool

func newBool(p *bool) *boolFlag {
	return (*boolFlag)(p)
}

func (b *boolFlag) Set(s string) error {
	val, err := strconv.ParseBool(s)
	*b = boolFlag(val)
	return err
}

func (b *boolFlag) String() string {
	return strconv.FormatBool(bool(*b))
}

func (b *boolFlag) Get() interface{} {
	return bool(*b)
}

type intFlag int

func newInt(def int, p *int) *intFlag {
	*p = def
	return (*intFlag)(p)
}

func (i *intFlag) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	*i = intFlag(v)
	return err
}

func (i *intFlag) Get() interface{} {
	return int(*i)
}

func (i *intFlag) String() string {
	return strconv.Itoa(int(*i))
}

type stringFlag string

func newString(def string, p *string) *stringFlag {
	*p = def
	return (*stringFlag)(p)
}

func (s *stringFlag) Set(t string) error {
	*s = stringFlag(t)
	return nil
}

func (s *stringFlag) Get() interface{} {
	return string(*s)
}

func (s *stringFlag) String() string {
	return string(*s)
}

type uintFlag uint

func newUint(def uint, p *uint) *uintFlag {
	*p = def
	return (*uintFlag)(p)
}

func (i *uintFlag) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	*i = uintFlag(v)
	return err
}

func (i *uintFlag) Get() interface{} {
	return uint(*i)
}

func (i *uintFlag) String() string {
	return strconv.FormatUint(uint64(*i), 10)
}
