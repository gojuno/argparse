package argparse

import (
	"bytes"
	"fmt"
	"log"
	// "strings"
)

// Parser commandline arguments parser
type Parser struct {
	items          map[string]CliItemType
	options        map[string]*CliItem
	arguments      []*CliItem
	unknownArgvLen bool
}

// InitParser parser
func ArgumentParser() (*Parser, error) {
	p := new(Parser)
	p.items = map[string]CliItemType{}
	p.options = map[string]*CliItem{}
	p.arguments = []*CliItem{}
	p.unknownArgvLen = false
	return p, nil
}

func Init(items []*CliItem) (*Parser, error) {
	p := new(Parser)
	p.items = map[string]CliItemType{}
	p.options = map[string]*CliItem{}
	p.arguments = []*CliItem{}
	p.unknownArgvLen = false
	err := p.Init(items)
	return p, err
}

func (p *Parser) String() string {
	buffer := bytes.NewBuffer([]byte{})
	for _, option := range p.options {
		fmt.Fprintf(buffer, "%v=%v\n", option.Name(), option)
	}
	for _, arg := range p.arguments {
		fmt.Fprintf(buffer, "%v=%v\n", arg.Name(), arg)
	}
	return buffer.String()
}

func (p *Parser) Dump() {
	log.Printf("%v", p.String())
}

func Option(name string) *CliItem {
	return NewCliItem(name, OPTION)
}

func Argument(name string) *CliItem {
	return NewCliItem(name, ARGUMENT)
}

// AddOption add option
func (p *Parser) Add(ci *CliItem) error {
	_, ok := p.items[ci.Name()]
	if ok {
		return fmt.Errorf("duplicate item [%s] defenition", ci.Name())
	}
	switch ci.Type() {
	case ARGUMENT:
		if p.unknownArgvLen {
			return fmt.Errorf("unknown argument [%v] values count", ci.Name())
		}
		switch ci.Param(NARG).Str() {
		case "1":
		case "*", "+":
			p.unknownArgvLen = true
		}
		p.arguments = append(p.arguments, ci)
	case OPTION:
		p.options[ci.Name()] = ci
	default:
		return fmt.Errorf("not implemented type [%s]", ci.Type())
	}

	p.items[ci.Name()] = ci.Type()
	return nil
}

// InitParser parser
func (p *Parser) Init(items []*CliItem) error {
	for _, ci := range items {
		if err := p.Add(ci); err != nil {
			return err
		}
	}
	return nil
}

/*
// AddStringCliItem add option
func (p *Parser) AddStringCliItem(name string, short string, long string) *CliItem {
	return p.AddCliItem(ARG_STRING, name)
}

// AddStringListCliItem add option
func (p *Parser) AddStringListCliItem(name string, short string, long string) *CliItem {
	return p.AddCliItem(ARG_STRING_LIST, name, short, long)
}

// AddFlagCliItem add option
func (p *Parser) AddFlagCliItem(name string, short string, long string) *CliItem {
	return p.AddCliItem(ARG_FLAG, name, short, long).Default("false").Action(SET_TRUE)
}

func (p *Parser) AddEnv(name string) *CliItem {
	option := NewCliItem(name)
	p.options[name] = option
	return option
}

func (p *Parser) Check(data string) (*CliItem, string) {
	// var option *CliItem
	// fmt.Printf("%v %v", p, data)
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
*/
func (p *Parser) parse(argv []string, ctx *ParserContext) error {

	longName := map[string]*CliItem{}
	shortName := map[string]*CliItem{}
	//argsIter := -1

	for _, item := range p.options {
		if short := item.Param(SHORT).Str(); short != "" {
			shortName[short] = item
		}
		if long := item.Param(LONG).Str(); long != "" {
			longName[long] = item
		}
	}

	// p.initDefaults(args)

	// p.context = NewParserContext(argv)
	// value, err := p.context.Next()
	// for err == nil {
	// 	o, v := p.Check(value)
	// 	if o == nil {
	// 		return fmt.Errorf("Unknown argument [%s]", value)
	// 	}
	// 	//log.Printf("! [%v] checked as %v\n", value, o)

	// 	switch o.optionType {
	// 	case ARG_FLAG:
	// 		switch o.action {
	// 		case SET_TRUE:
	// 			value = "true"
	// 		case SET_FALSE:
	// 			value = "false"
	// 		default:
	// 			value = o.defaultValue
	// 		}
	// 	case ARG_STRING, ARG_STRING_LIST:
	// 		if v != "" {
	// 			value = v
	// 		} else {
	// 			value, err = p.context.Next()
	// 			if err != nil {
	// 				return fmt.Errorf("Argument [%s] value required", o)
	// 			}
	// 		}
	// 	case ARG_ARGS:
	// 		switch o.narg {
	// 		case "1":
	// 			value, err = p.context.Next()
	// 		case "+":
	// 			for err == nil {
	// 				value, err = p.context.Next()
	// 			}
	// 		case "*":
	// 			for {
	// 				value, err = p.context.Next()
	// 			}
	// 		default:
	// 			return fmt.Errorf("Incorrect narg %v", o)
	// 		}
	// 	}

	// 	args.Save(o.name, o.optionType, value)

	// 	value, err = p.context.Next()
	// }
	// if err.Error() != EOF {
	// 	return err
	// }
	//	return p.checkRequired(args)
	return nil
}

/*
// Parser commandline arguments parser
type Parser struct {
	options map[string]*CliItem
}

// InitParser parser
func EnvParser() (*Parser, error) {
	p := new(Parser)
	p.options = map[string]*CliItem{}
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
