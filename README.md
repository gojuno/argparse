# argparse

command-line parsing module

## Example

```
package main

import (
	"fmt"

	"github.com/juno-lab/argparse"
)

func main() {
	parser, _ := argparse.ArgumentParser()
	parser.AddOption("input", "i", "input", false)
	parser.AddOption("output", "o", "output", false)
	parser.AddOption("libs", "L", "", true)

	args := parser.ParseArgs()
	fmt.Fprintf(os.Stderr, "%v\n", args)
	v, ok := args.AsList("libs")
	fmt.Fprintf(os.Stderr, "%v %v\n", ok, v)
	v1, ok1 := args.AsString("input")
	fmt.Fprintf(os.Stderr, "%v %v\n", ok1, v1)
	v2, ok2 := args.AsString("output")
	fmt.Fprintf(os.Stderr, "%v %v\n", ok2, v2)
}
```
