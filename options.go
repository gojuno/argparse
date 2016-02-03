package argparse

import (
	"bytes"
	"fmt"
)

type OptionAction string

const (
	DEFAULT_ACTION OptionAction = ""
	SET_TRUE       OptionAction = "SET_TRUE"
	SET_FALSE      OptionAction = "SET_FALSE"
)

// type CmdOption interface {
//     Type()
//     Name()
//     Short()
//     Long()
// }

// Option option
type Option struct {
	optionType   ArgumentType
	name         string
	short        string
	long         string
	defaultValue string
	action       OptionAction
	narg         string
	required     bool
}

func NewOption(name string) *Option {
	option := new(Option)
	option.name = name
	option.defaultValue = ""
	option.required = false
	return option
}

// String string
func (o *Option) String() string {
	buffer := bytes.NewBuffer([]byte{})
	fmt.Fprintf(buffer, "%v", o.name)
	if o.short != "" && o.long != "" {
		fmt.Fprintf(buffer, "(-%v|--%v)", o.short, o.long)
	} else if o.short != "" {
		fmt.Fprintf(buffer, "(-%v)", o.short)
	} else {
		fmt.Fprintf(buffer, "(--%v)", o.long)
	}
	return buffer.String()
}

// Name name
func (o *Option) Name() string {
	return o.name
}

// Short name
func (o *Option) Short() string {
	return o.short
}

// Long name
func (o *Option) Long() string {
	return o.long
}

// Default default
func (o *Option) Default(defaultValue string) *Option {
	o.defaultValue = defaultValue
	o.required = false
	return o
}

// Action default
func (o *Option) Action(action OptionAction) *Option {
	o.action = action
	return o
}

// NArg default
func (o *Option) NArg(narg string) *Option {
	o.narg = narg
	switch narg {
	case "1", "+":
		o.required = true
	case "*":
		o.required = false
	}
	return o
}

// Required required
func (o *Option) Required() *Option {
	o.required = true
	return o
}
