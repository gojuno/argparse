package argparse

import (
	// "fmt"
	// "os"
	"testing"
)

func TestCli(t *testing.T) {
	// t.Error("Expect logs to contain escaped secret field:\n", logs)
	items := []*CliItem{
		NewCliItem("name", OPTION),
		NewCliItem("name", OPTION).Required(),
		NewCliItem("name", OPTION).Default("default"),
	}

	expectedResults := []string{
		"[name]: <short=<nil>;long=name;default=<nil>;action=<nil>;required=false;narg=<nil>;>%!(EXTRA string=name)",
		"[name]: <short=<nil>;long=name;default=<nil>;action=<nil>;required=true;narg=<nil>;>%!(EXTRA string=name)",
		"[name]: <short=<nil>;long=name;default=default;action=<nil>;required=false;narg=<nil>;>%!(EXTRA string=name)",
	}
	for i, ci := range items {
		if expected := expectedResults[i]; ci.String() != expected {
			t.Errorf("cliItem [%v] error. Expected [%v], got [%v]", ci.Name(), expected, ci.String())
		}
	}
}
