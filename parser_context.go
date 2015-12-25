package argparse

import (
	"errors"
)

const (
	EOF = "EOF"
)

// CmdArgvParserContext commandline arguments parser
type ParserContext struct {
	rawArgv []string
	iter    int
}

func NewParserContext(argv []string) *ParserContext {
	pc := new(ParserContext)
	pc.iter = -1
	pc.rawArgv = []string{}
	for _, d := range argv {
		pc.rawArgv = append(pc.rawArgv, d)
	}
	return pc
}

// Next get string value
func (pc *ParserContext) Next() (string, error) {
	if pc.iter == len(pc.rawArgv)-1 {
		return "", errors.New(EOF)
	}
	pc.iter++
	return pc.rawArgv[pc.iter], nil
}
