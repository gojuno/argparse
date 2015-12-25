package argparse

import (
	// "fmt"
	// "os"
	"testing"
)

func TestParserStringOptions(t *testing.T) {
	// t.Error("Expect logs to contain escaped secret field:\n", logs)

	parser, _ := ArgumentParser()
	parser.AddStringOption("s1", "a", "") // -a value
	parser.AddStringOption("s2", "b", "") // -avalue
	parser.AddStringOption("s3", "", "c") // --arg value
	parser.AddStringOption("s4", "", "d") // --arg=value
	parser.AddStringOption("s5", "", "e").Default("s5")
	parser.AddStringOption("s6", "", "f").Default("default")

	argv := []string{
		"-a", "s1",
		"-bs2",
		"--c", "s3",
		"--d=s4",
		"--f=s6",
	}
	// args := parser.ParseArgs()
	args := NewArgs()

	if err := parser.parse(argv, args); err != nil {
		t.Error("Parse error:\n", argv, err)
		t.SkipNow()
	}
	for i, v := range []string{"s1", "s2", "s3", "s4", "s5", "s6"} {
		a := args.Arg(v)
		if a != v {
			t.Errorf("[%02d] arg [%v] error. Expected [%s], got [%s]", i, v, v, a)
		}
	}
}

func TestParserStringListOptions(t *testing.T) {
	// t.Error("Expect logs to contain escaped secret field:\n", logs)

	parser, _ := ArgumentParser()

	parser.AddStringListOption("l1", "a", "") // -a value
	parser.AddStringListOption("l2", "b", "") // -avalue
	parser.AddStringListOption("l3", "", "c") // --arg value
	parser.AddStringListOption("l4", "", "d") // --arg=value

	argv := []string{
		"-a", "l1", "-a", "l2",
		"-b", "l3",
		"--c", "l3",
		"--d", "l4",
	}
	// args := parser.ParseArgs()
	args := NewArgs()

	if err := parser.parse(argv, args); err != nil {
		t.Error("Parse error:\n", argv, err)
		t.SkipNow()
	}
	a := args.Arg("l1").([]string)
	if len(a) != 2 || a[0] != "l1" || a[1] != "l2" {
		t.Errorf("arg 'l1' error. Expected [l1, l2], got [%v]", a)
	}
	a = args.Arg("l2").([]string)
	if len(a) != 1 || a[0] != "l3" {
		t.Errorf("arg 'l2' error. Expected [l2], got [%v]", a)
	}
	a = args.Arg("l3").([]string)
	if len(a) != 1 || a[0] != "l3" {
		t.Errorf("arg 'l3' error. Expected [true], got [%v]", a)
	}
	a = args.Arg("l4").([]string)
	if len(a) != 1 || a[0] != "l4" {
		t.Errorf("arg 'l4' error. Expected [true], got [%v]", a)
	}

}

func TestParserFlagOptions(t *testing.T) {
	// t.Error("Expect logs to contain escaped secret field:\n", logs)

	parser, _ := ArgumentParser()

	parser.AddFlagOption("f1", "a", "")
	parser.AddFlagOption("f2", "", "b")
	parser.AddFlagOption("f3", "c", "").Default("true").Action(SET_FALSE)
	parser.AddFlagOption("f4", "", "d").Default("true").Action(SET_FALSE)

	argv := []string{
		"-a",
		"--b",
		// "-c",
		"--d",
	}
	// args := parser.ParseArgs()
	args := NewArgs()

	if err := parser.parse(argv, args); err != nil {
		t.Error("Parse error:\n", argv, err)
		t.SkipNow()
	}

	if a := args.Arg("f1"); a != true {
		t.Errorf("arg 'f1' error. Expected [true], got [%v]", a)
	}
	if a := args.Arg("f2"); a != true {
		t.Errorf("arg 'f2' error. Expected [true], got [%v]", a)
	}
	if a := args.Arg("f3"); a != true {
		t.Errorf("arg 'f3' error. Expected [true], got [%v]", a)
	}
	if a := args.Arg("f4"); a != true {
		t.Errorf("arg 'f4' error. Expected [true], got [%v]", a)
	}
}

func TestParserOptions(t *testing.T) {
	t.SkipNow()
	// t.Error("Expect logs to contain escaped secret field:\n", logs)

	parser, _ := ArgumentParser()
	parser.AddStringOption("string_option_short", "s", "")
	parser.AddStringOption("string_option_long", "", "string-option")
	parser.AddStringOption("string_option_default", "d", "").Default("default")

	parser.AddStringListOption("string_list_option", "l", "string-list-option")

	parser.AddFlagOption("flag_option", "f", "")
	parser.AddFlagOption("flag_option_default", "", "flag-option").Default("true")
	parser.AddFlagOption("flag_option_action", "a", "").Action(SET_TRUE)

	// parser.AddStringOption("string_option_with_default", "s", "string-option-with-default").Default("default")
	// parser.AddStringOption("output", "o", "output").Default("<STDOUT>")
	// parser.AddListOption("libs", "L", "").List()
	// parser.AddBoolOption("no_pretty", "", "no_pretty").Default(false).Action(SET_TRUE)

	argv := []string{
		"-s", "data1",
		"-l", "data2",
		"-ldata3",
		//"-f",

		"--string-option", "data4",
		"--string-list-option", "data5",
		"--string-list-option=data6",
		"--flag-option",
		"-a",
	}
	// args := parser.ParseArgs()
	args := NewArgs()

	if err := parser.parse(argv, args); err != nil {
		t.Error("Parse error:\n", argv, err)
		t.SkipNow()
	}
}
