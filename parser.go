package argparse

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	EOF = "EOF"
)

// CmdArgvParser commandline arguments parser
type CmdArgvParser struct {
	rawArgv []string
	iter    int
	options map[string]*CmdOption
}

// InitParser parser
func ArgumentParser() (*CmdArgvParser, error) {
	p := new(CmdArgvParser)
	p.options = map[string]*CmdOption{}
	p.iter = -1

	for _, v := range os.Args[1:] {
		p.rawArgv = append(p.rawArgv, v)
	}
	return p, nil
}

// Next get string values
func (p *CmdArgvParser) AddArgument() {
}

// AddOption add option
func (p *CmdArgvParser) AddOption(name string, short string, long string, list bool) {
	option := &CmdOption{
		name:  name,
		short: "-" + short,
		long:  "--" + long,
		list:  list,
	}
	p.options[name] = option
}

// Next get string values
func (p *CmdArgvParser) Next() (string, error) {
	if p.iter == len(p.rawArgv)-1 {
		return "", errors.New(EOF)
	}
	p.iter++
	return p.rawArgv[p.iter], nil
}

func (p *CmdArgvParser) Save(args *CmdArgs, option *CmdOption, value string) {
	if option.list {
		args.Add(option.name, value)
	} else {
		args.Set(option.name, value)
	}
}

func (p *CmdArgvParser) parse(args *CmdArgs) error {
	value, err := p.Next()
	for err == nil {
		for _, o := range p.options {
			// fmt.Fprintf(os.Stderr, "* checking [%v] %v\n", value, o)
			if !o.Check(value) {
				continue
			}
			// fmt.Fprintf(os.Stderr, "! checking [%v] %v\n", value, o)

			if o.short != "" { // -o value | -ovalue
				if value == o.short {
					if v, e := p.Next(); e != nil {
						return errors.New("short option processing error")
					} else {
						p.Save(args, o, v)
					}
				} else if strings.HasPrefix(value, o.short) {
					p.Save(args, o, value[len(o.short):])
				} else {
					return errors.New("short option processing error")
				}
			} else if o.long != "" && strings.HasPrefix(value, o.long) { // --opt value | --opt=value
				if value == o.long {
					if v, e := p.Next(); e != nil {
						return errors.New("long option processing error")
					} else {
						p.Save(args, o, v)
					}
				} else if strings.HasPrefix(value, o.long+"=") {
					p.Save(args, o, value[len(o.long):])
				} else {
					return errors.New("long option processing error")
				}
			}
		}
		value, err = p.Next()
	}
	if err.Error() != EOF {
		return err
	}
	return nil
}

// ParseArgs process command line
func (p *CmdArgvParser) ParseArgs() *CmdArgs {
	args := &CmdArgs{}
	args.Init()
	// args.Set("input", "<STDIN>")
	// args.Set("output", "<STDOUT>")

	if err := p.parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "processing error: %v\n", err)
		os.Exit(1)
	}

	return args
}
