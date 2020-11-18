package flagger

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrNoFlags is the error returned when no flags are passed.
	// Useful as a check for when you want to proceed without flags.
	ErrNoFlags = errors.New("No flags passed")
	ErrHelp    = errors.New("Help")
	ErrVersion = errors.New("Version")
)

// Flags is the struct used by you, the user
type Flags struct {
	flags       []*Flag
	available   map[string]int
	help        bool
	helpText    string
	version     bool
	versionText string
}

// Bool creates a bool flag, returns the bool
// usage is the string for the help
// flgs
func (f *Flags) Bool(usage string, flgs ...string) *bool {
	p := new(bool)
	f.add(&Flag{flags: flgs, Usage: usage, Value: newBool(p)})
	return p
}

func (f *Flags) BoolVar(b *bool, usage string, flgs ...string) {
	f.add(&Flag{flags: flgs, Usage: usage, Value: newBool(b)})
}

func (f *Flags) Int(def int, usage string, flgs ...string) *int {
	p := new(int)
	f.add(&Flag{flags: flgs, Usage: usage, Value: newInt(def, p)})
	return p
}

func (f *Flags) IntVar(i *int, def int, usage string, flgs ...string) {
	f.add(&Flag{flags: flgs, Usage: usage, Value: newInt(def, i)})
}

func (f *Flags) String(def string, usage string, flgs ...string) *string {
	p := new(string)
	f.add(&Flag{flags: flgs, Usage: usage, Value: newString(def, p)})
	return p
}

func (f *Flags) StringVar(s *string, def string, usage string, flgs ...string) {
	f.add(&Flag{flags: flgs, Usage: usage, Value: newString(def, s)})

}

func (f *Flags) Uint(def uint, usage string, flgs ...string) *uint {
	p := new(uint)
	f.add(&Flag{flags: flgs, Usage: usage, Value: newUint(def, p)})
	return p
}

func (f *Flags) UintVar(i *uint, def uint, usage string, flgs ...string) {
	f.add(&Flag{flags: flgs, Usage: usage, Value: newUint(def, i)})
}

func (f *Flags) add(flg *Flag) {
	index := len(f.flags)
	f.flags = append(f.flags, flg)
	for i := range flg.flags {
		f.available[flg.flags[i]] = index
	}
}

func (f Flags) Help(msg string) {
	fmt.Println(msg)
	for i := range f.flags {
		fmt.Println(f.flags[i].Print())
	}
}

func sanitize(s []string) []string {
	a := make([]string, 0)
	for i := range s {
		if len(s[i]) > 0 {
			if s[i][:1] == "-" {
				if len(s[i]) < 2 {
					continue
				}
				if s[i][:2] == "--" {
					q := strings.Split(s[i], "=")
					a = append(a, q[0])
					if len(q) > 1 {
						z := ""
						x := q[1:]
						for k := range x {
							if z == "" {
								z = x[k]
							} else {
								z = z + "=" + x[k]
							}
						}
						a = append(a, z)
					}
				} else {
					if len(s[i]) == 2 {
						a = append(a, s[i])
					} else {
						a = append(a, s[i][:2])
						temp := s[i][2:]
						x := 0
						for y := x + 1; y <= len(temp); y++ {
							if temp[x:y] == "=" {
								a = append(a, temp[y:])
								break
							} else {
								a = append(a, "-"+temp[x:y])
							}
							x++
						}
					}
				}
			} else {
				a = append(a, s[i])
			}
		}
	}
	return a
}

func (f *Flags) Parse(flags []string) ([]string, error) {
	var err error
	if len(flags) < 1 {
		return nil, ErrNoFlags
	}
	data := make([]string, 0)
	fgs := sanitize(flags)
	var found *Flag
	for i := 0; i < len(fgs); i++ {
		if fgs[i][:1] != "-" {
			data = append(data, fgs[i])
			continue
		} else {
			if index, ok := f.available[fgs[i]]; ok {
				found = f.flags[index]
			} else {
				return nil, fmt.Errorf("No such flag %s", fgs[i])
			}
		}
		switch found.Value.(type) {
		case *boolFlag:
			if err = found.Value.Set("true"); err != nil {
				return nil, err
			}
		default:
			if (i + 1) < len(fgs) {
				i++
				if err = found.Value.Set(fgs[i]); err != nil {
					return nil, fmt.Errorf("Value %s is not compatible with flag %s", fgs[i], fgs[i-1])
				}
			} else {
				return nil, fmt.Errorf("No value passed into flag %s", fgs[i])
			}
		}

	}
	if err == nil && f.help {
		f.Help(f.helpText)
		return nil, ErrHelp
	}
	if err == nil && f.version {
		fmt.Println(f.versionText)
		return nil, ErrVersion
	}
	return data, err
}

func (f *Flags) AddHelp(txt string, msg string) error {
	f.helpText = msg
	f.BoolVar(&f.help, txt, "-h", "--help")
	return nil
}

func (f *Flags) AddVersion(txt string, msg string) error {
	f.versionText = msg
	f.BoolVar(&f.version, txt, "-v", "--version")
	return nil
}

func New() *Flags {
	f := &Flags{}
	f.flags = make([]*Flag, 0)
	f.available = make(map[string]int)

	return f
}
