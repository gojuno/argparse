package argparse

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

// Parser commandline arguments parser
type Parser struct {
	options    map[string]*Option
	shortNames map[string]*Option
	longNames  map[string]*Option

	context *ParserContext

	optionsProcessed bool
	args             []*Option
	argsIdx          int
}

// InitParser parser
func ArgumentParser() (*Parser, error) {
	p := new(Parser)
	p.optionsProcessed = false
	p.args = []*Option{}
	p.argsIdx = -1
	p.options = map[string]*Option{}
	p.shortNames = map[string]*Option{}
	p.longNames = map[string]*Option{}
	return p, nil
}

func (p *Parser) String() string {
	buffer := bytes.NewBuffer([]byte{})
	for name, option := range p.options {
		fmt.Fprintf(buffer, "%v=%v\n", name, option)
	}
	return buffer.String()
}

func (p *Parser) Dump() {
	log.Printf("%v", p.String())
}

func (p *Parser) Check(data string) (*Option, string) {
	// var option *Option
	// log.Printf("%v", p.String())
	if !p.optionsProcessed {
		if strings.HasPrefix(data, "--") { // --arg | --arg value | --arg=value
			argData := data[2:]
			if o, ok := p.longNames[argData]; ok {
				return o, ""
			} else if idx := strings.Index(argData, "="); idx != -1 {
				if o, ok = p.longNames[argData[:idx]]; ok {
					return o, argData[idx+1:]
				}
			}
		} else if strings.HasPrefix(data, "-") { // -a | -avalue | -a value
			argData := data[1:]
			if o, ok := p.shortNames[argData]; ok {
				return o, ""
			} else if o, ok := p.shortNames[argData[:1]]; ok {
				return o, argData[1:]
			}
		} else { // first arg
			if !p.optionsProcessed {
				p.optionsProcessed = true
			}
			p.argsIdx++
			o := p.args[p.argsIdx]
			return o, ""
		}
	} else {
		if strings.HasPrefix(data, "-") {
			// error. option beetwen args
			return nil, ""
		}
	}
	return nil, ""
}

// AddOption add option
func (p *Parser) AddOption(optionType ArgumentType, name string, short string, long string) *Option {
	option := NewOption(name)
	option.short = short
	option.long = long
	option.optionType = optionType

	p.options[name] = option
	if short != "" {
		p.shortNames[short] = option
	}
	if long != "" {
		p.longNames[long] = option
	}
	return option
}

// AddArg add option
func (p *Parser) AddArg(name string) *Option {
	return p.AddOption(ARG_ARGS, name, "", "").NArg("1")
}

// AddStringOption add option
func (p *Parser) AddStringOption(name string, short string, long string) *Option {
	return p.AddOption(ARG_STRING, name, short, long)
}

// AddStringListOption add option
func (p *Parser) AddStringListOption(name string, short string, long string) *Option {
	return p.AddOption(ARG_STRING_LIST, name, short, long)
}

// AddFlagOption add option
func (p *Parser) AddFlagOption(name string, short string, long string) *Option {
	return p.AddOption(ARG_FLAG, name, short, long).Default("false").Action(SET_TRUE)
}

func (p *Parser) AddEnv(name string) *Option {
	option := NewOption(name)
	p.options[name] = option
	return option
}

func (p *Parser) Parse() *Args {
	args := NewArgs()
	if err := p.parse(os.Args[1:], args); err != nil {
		log.Fatalf("ParseArgs: %v\n", err)
	}
	return args
}

func (p *Parser) initDefaults(args *Args) {
	for name, option := range p.options {
		switch option.optionType {
		case ARG_FLAG, ARG_STRING:
			if option.defaultValue != "" {
				args.Save(name, option.optionType, option.defaultValue)
			}
		}
	}
}

func (p *Parser) checkRequired(args *Args) error {
	for name, option := range p.options {
		if args.Arg(name) == nil {
			return fmt.Errorf("Field [%v] required", option)
		}
	}
	return nil
}

func (p *Parser) parse(argv []string, args *Args) error {
	p.initDefaults(args)

	p.context = NewParserContext(argv)
	value, err := p.context.Next()
	for err == nil {
		o, v := p.Check(value)
		if o == nil {
			return fmt.Errorf("Unknown argument [%s]", value)
		}
		//log.Printf("! [%v] checked as %v\n", value, o)

		switch o.optionType {
		case ARG_FLAG:
			switch o.action {
			case SET_TRUE:
				value = "true"
			case SET_FALSE:
				value = "false"
			default:
				value = o.defaultValue
			}
		case ARG_STRING, ARG_STRING_LIST:
			if v != "" {
				value = v
			} else {
				value, err = p.context.Next()
				if err != nil {
					return fmt.Errorf("Argument [%s] value required", o)
				}
			}
		case ARG_ARGS:
			switch o.narg {
			case "1":
				value, err = p.context.Next()
			case "+":
				for err == nil {
					value, err = p.context.Next()
				}
			case "*":
				for {
					value, err = p.context.Next()
				}
			default:
				return fmt.Errorf("Incorrect narg %v", o)
			}
		}

		args.Save(o.name, o.optionType, value)

		value, err = p.context.Next()
	}
	if err.Error() != EOF {
		return err
	}
	return p.checkRequired(args)
}

/*
// Parser commandline arguments parser
type Parser struct {
	options map[string]*Option
}

// InitParser parser
func EnvParser() (*Parser, error) {
	p := new(Parser)
	p.options = map[string]*Option{}
	return p, nil
}

func (p *Parser) Init() {
	fmt.Println(os.Environ())
}

func (p *Parser) initEnvironment() *Environment {
	env := NewEnv()
	fmt.Println(os.Environ())
	return env
}
*/
