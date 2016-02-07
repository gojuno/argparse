package argparse

import (
	"bytes"
	"fmt"
)

type CliItemType string

const (
	OPTION   CliItemType = "option"
	ARGUMENT CliItemType = "argument"
)

type CliItemAction string

const (
	DEFAULT_ACTION CliItemAction = ""
	SET_TRUE       CliItemAction = "SET_TRUE"
	SET_FALSE      CliItemAction = "SET_FALSE"
)

type CliItemParameterType string

const (
	SHORT    CliItemParameterType = "short"
	LONG     CliItemParameterType = "long"
	DEFAULT  CliItemParameterType = "default"
	ACTION   CliItemParameterType = "action"
	REQUIRED CliItemParameterType = "required"

	LIST     CliItemParameterType = "list"
	CSV_LIST CliItemParameterType = "csv_list"

	NARG CliItemParameterType = "narg"
)

// CliItem CliItem
type CliItem struct {
	name   string
	ciType CliItemType
	params map[CliItemParameterType]CliItemParam
}

type CliItemParam struct {
	value interface{}
}

func (p CliItemParam) Int() int {
	return p.value.(int)
}

func (p CliItemParam) Str() string {
	return p.value.(string)
}

func (p CliItemParam) Bool() bool {
	return p.value.(bool)
}

func NewCliItem(name string, ciType CliItemType) *CliItem {
	item := new(CliItem)
	item.name = name
	item.ciType = ciType
	item.params = map[CliItemParameterType]CliItemParam{}
	item.SetParam(REQUIRED, false)
	item.SetParam(LONG, name)
	return item
}

// String string
func (ci *CliItem) String() string {
	buffer := bytes.NewBuffer([]byte{})
	fmt.Fprintf(buffer, "[%v]: <", ci.name)
	for _, i := range []CliItemParameterType{
		SHORT,
		LONG,
		DEFAULT,
		ACTION,
		REQUIRED,
		NARG,
	} {
		fmt.Fprintf(buffer, "%v=%v;", i, ci.Param(i).value)
	}
	fmt.Fprintf(buffer, ">", ci.name)
	return buffer.String()
}

// Name name
func (ci *CliItem) Name() string {
	return ci.name
}

// Type type
func (ci *CliItem) Type() CliItemType {
	return ci.ciType
}

// Param name
func (ci *CliItem) Param(t CliItemParameterType) CliItemParam {
	return ci.params[t]
}

// Param name
func (ci *CliItem) SetParam(t CliItemParameterType, v interface{}) *CliItem {
	ci.params[t] = CliItemParam{value: v}
	return ci
}

// Short name
func (ci *CliItem) Short(s string) *CliItem {
	return ci.SetParam(SHORT, s)
}

// Long name
func (ci *CliItem) Long(s string) *CliItem {
	return ci.SetParam(LONG, s)
}

// Default default
func (ci *CliItem) Default(defaultValue string) *CliItem {
	return ci.SetParam(DEFAULT, defaultValue).SetParam(REQUIRED, false)
}

// Required required
func (ci *CliItem) Required() *CliItem {
	return ci.SetParam(REQUIRED, true)
}

// Action default
func (ci *CliItem) Action(action CliItemAction) *CliItem {
	return ci.SetParam(ACTION, action)
}

// List default
func (ci *CliItem) List() *CliItem {
	return ci.SetParam(LIST, true)
}

// List default
func (ci *CliItem) CsvList() *CliItem {
	return ci.SetParam(CSV_LIST, true)
}

// NArg default
func (ci *CliItem) NArg(narg string) *CliItem {
	ci.SetParam(NARG, narg)
	switch narg {
	case "1", "+":
		ci.SetParam(REQUIRED, true)
	case "*":
		ci.SetParam(REQUIRED, false)
	}
	return ci
}
