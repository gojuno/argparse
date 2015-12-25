package argparse

import (
	"bytes"
	"fmt"
	"log"
)

// CmdArg commandline argument
type Args struct {
	argv map[string]Argument
}

// String get string value
func (a *Args) String() string {
	buffer := bytes.NewBuffer([]byte{})
	for name, arg := range a.argv {
		fmt.Fprintf(buffer, "%v=%v\n", name, arg)
	}
	return buffer.String()
}

func NewArgs() *Args {
	args := new(Args)
	args.argv = map[string]Argument{}
	return args
}

func (a *Args) Arg(name string) interface{} {
	v, ok := a.argv[name]
	if !ok {
		return nil
	}
	return v.Value()
}

func (a *Args) AsString(name string) string {
	str, _ := a.Arg(name).(string)
	return str
}

func (a *Args) AsList(name string) []string {
	lst, _ := a.Arg(name).([]string)
	return lst
}

func (a *Args) AsFlag(name string) bool {
	flg, _ := a.Arg(name).(bool)
	return flg
}

func (a *Args) Save(name string, optionType ArgumentType, value string) {
	var argument Argument
	if arg, ok := a.argv[name]; !ok {
		switch optionType {
		case ARG_FLAG:
			argument = new(FlagArg)
		case ARG_STRING:
			argument = new(StringArg)
		case ARG_STRING_LIST:
			argument = new(StringListArg)
		default:
			log.Printf("! ERROR\n")
		}
		a.argv[name] = argument
	} else {
		argument = arg
	}
	switch optionType {
	case ARG_FLAG, ARG_STRING:
		argument.Set(value)
	case ARG_STRING_LIST:
		argument.Add(value)
	default:
		log.Printf("! ERROR %v %v = %v\n", optionType, name, value)
	}
	// if option.list {
	// 	args.Add(option, value)
	// } else {
	// 	args.Set(option, value)
	// }
}

/*
import (
	// "errors"
	"bytes"
	"fmt"
	// "strings"
)

// CmdArg commandline argument
type CmdArg struct {
	option *CmdOption
	raw    []string
}


// AsString get string value
func (arg CmdArg) AsString() string {
	return arg.raw[0]
}

// AsBool get bool value
func (arg CmdArg) AsBool() bool {
	return arg.raw[0] == "true" || arg.raw[0] == "True" || arg.raw[0] == "1"
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
func (args *CmdArgs) Init(options map[string]*CmdOption) {
	args.argv = map[string]*CmdArg{}
	for _, option := range options {
		args.argv[option.name] = &CmdArg{option: option, raw: []string{""}}
	}
}

func (args *CmdArgs) Info() string {
	buffer := bytes.NewBuffer([]byte{})
	for name, arg := range args.argv {
		fmt.Fprintf(buffer, "%v=%v\n", name, arg)
		switch {
		case !arg.option.list && arg.option.action == DEFAULT_ACTION:
			fmt.Fprintf(buffer, "%v=%v\n", name, arg.AsString())
		case !arg.option.list && arg.option.action != DEFAULT_ACTION:
			fmt.Fprintf(buffer, "%v=%v\n", name, arg.AsBool())
		case arg.option.list:
			fmt.Fprintf(buffer, "%v=%v\n", name, arg.AsStrings())
		default:
			fmt.Fprintf(buffer, "unknown arg type\n")
		}
	}
	return buffer.String()
}

// Set value
func (args *CmdArgs) Set(option *CmdOption, value string) {
	arg, ok := args.argv[option.name]
	if !ok {
	}
	arg.raw[0] = value
}

// Add value
func (args *CmdArgs) Add(option *CmdOption, value string) {
	arg, ok := args.argv[option.name]
	if !ok {
	}
	arg.raw = append(arg.raw, value)
}

// AsBool get bool value
func (args *CmdArgs) AsBool(name string) (bool, bool) {
	result := ""
	r, ok := args.argv[name]
	if len(r.raw) > 0 && r.raw[0] != "" {
		result = r.raw[0]
	} else if r.option.defaultValue != nil {
		result = r.option.defaultValue.(string)
	}
	return result == "true" || result == "True" || result == "1", ok
}

// AsString get string value
func (args *CmdArgs) AsString(name string) (string, bool) {
	result := ""
	r, ok := args.argv[name]
	if len(r.raw) > 0 && r.raw[0] != "" {
		result = r.raw[0]
	} else if r.option.defaultValue != nil {
		result = r.option.defaultValue.(string)
	}
	return result, ok
}

// AsList get string value
func (args *CmdArgs) AsList(name string) ([]string, bool) {
	result := []string{}
	r, ok := args.argv[name]
	if ok {
		result = r.raw
	} else if r.option.defaultValue != nil {
		result = []string{r.option.defaultValue.(string)}
	}
	return result, ok
}
*/
