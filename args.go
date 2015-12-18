package argparse

import (
	// "errors"
	"bytes"
	"fmt"
	// "strings"
)

// CmdArg commandline argument
type CmdArg struct {
	name string
	raw  []string
}

// String get string value
func (arg CmdArg) String() string {
	return fmt.Sprintf("%v=%v", arg.name, arg.raw)
}

// AsString get string value
func (arg CmdArg) AsString() string {
	return arg.raw[0]
}

// AsStrings get string values
func (arg CmdArg) AsStrings() []string {
	return arg.raw
}

// CmdArgs commandline arguments
type CmdArgs struct {
	argv map[string]*CmdArg
}

// String get string value
func (args CmdArgs) String() string {
	buffer := bytes.NewBuffer([]byte{})
	for k, v := range args.argv {
		fmt.Fprintf(buffer, "%s %s\n", k, v)
	}
	return buffer.String()
}

// Init init
func (args *CmdArgs) Init() {
	args.argv = map[string]*CmdArg{}
}

// Set value
func (args *CmdArgs) Set(name string, value string) {
	_, ok := args.argv[name]
	if !ok {
		args.argv[name] = &CmdArg{name: name}
		args.argv[name].raw = []string{""}
	}
	args.argv[name].raw[0] = value
}

// Add value
func (args *CmdArgs) Add(name string, value string) {
	_, ok := args.argv[name]
	if !ok {
		args.argv[name] = &CmdArg{name: name}
		args.argv[name].raw = []string{}
	}
	args.argv[name].raw = append(args.argv[name].raw, value)
}

// String get string value
func (args *CmdArgs) AsString(name string) (string, bool) {
	result := ""
	r, ok := args.argv[name]
	ok = ok && len(r.raw) > 0
	if ok {
		result = r.raw[0]
	}
	return result, ok
}

// String get string value
func (args *CmdArgs) AsList(name string) ([]string, bool) {
	result := []string{}
	r, ok := args.argv[name]
	if ok {
		result = r.raw
	}
	return result, ok
}
