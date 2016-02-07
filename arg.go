package argparse

import ()

/*
type ArgumentType string

const (
	ARG_STRING      ArgumentType = "ARG_STRING"
	ARG_STRING_LIST ArgumentType = "ARG_STRING_LIST"
	ARG_FLAG        ArgumentType = "ARG_FLAG"
	ARG_ARGS        ArgumentType = "ARG_ARGS"
)

type Argument interface {
	Value() interface{}
	Type() ArgumentType
	Add(string)
	Set(string)
}

type StringArg struct {
	value string
}

func (s *StringArg) Type() ArgumentType {
	return ARG_STRING
}

func (s *StringArg) Value() interface{} {
	return s.value
}

func (s *StringArg) Add(value string) {
	s.value = value
}

func (s *StringArg) Set(value string) {
	s.value = value
}

type StringListArg struct {
	value []string
}

func (s *StringListArg) Type() ArgumentType {
	return ARG_STRING_LIST
}

func (s *StringListArg) Value() interface{} {
	return s.value
}

func (s *StringListArg) Add(value string) {
	s.value = append(s.value, value)
}

func (s *StringListArg) Set(value string) {
	s.value = []string{value}
}

type FlagArg struct {
	value bool
}

func (s *FlagArg) Type() ArgumentType {
	return ARG_FLAG
}

func (s *FlagArg) Value() interface{} {
	return s.value
}

func (s *FlagArg) Add(value string) {
	s.value = value == "true" || value == "True" || value == "1"
}

func (s *FlagArg) Set(value string) {
	s.value = value == "true" || value == "True" || value == "1"
}
*/
