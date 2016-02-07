package argparse

import (
	"errors"
)

const (
	EOF = "EOF"
)

// CmdArgvParserContext commandline arguments parser
type ParserContext struct {
	argv []string
	iter int
}

func NewParserContext(argv []string) *ParserContext {
	pc := new(ParserContext)
	pc.argv = argv
	pc.iter = -1
	return pc
}

// Next extract next string value from input stream
func (pc *ParserContext) Next() (string, error) {
	if pc.iter == len(pc.argv)-1 {
		return "", errors.New(EOF)
	}
	pc.iter++
	return pc.argv[pc.iter], nil
}
