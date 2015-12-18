package argparse

import (
	"strings"
)

// CmdOption option
type CmdOption struct {
	name  string
	short string
	long  string
	list  bool
}

// Check option
func (o *CmdOption) Check(value string) bool {
	return (o.short != "" && strings.HasPrefix(value, o.short)) || (o.long != "" && value == o.long)
}
