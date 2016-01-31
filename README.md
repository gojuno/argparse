# argparse

command-line parsing module

support:
- short and long command line options (`-s | --long / -s<OPTION> | --long=<OPTION>`)
- command line arguments
- environment variables


## Example

```
ENV_VAR=variable your_command -v --check -L EN --lang=RU -LDE --libs=common,private,public dest.file.txt src_1.txt src_2.txt src_3.txt
```

source:

```
package main

import (
	"fmt"

	"github.com/juno-lab/argparse"
)

func main() {
	parser, _ := argparse.ArgumentParser()
	parser.AddFlagOption("verbose", "v", "verbose").Default("false").Action(SET_TRUE)
	parser.AddFlagOption("check", "", "check").Default("false").Action(SET_TRUE)

	parser.AddStringOption("lang", "L", "lang").List()
	parser.AddStringOption("libs", "", "libs").Csv()

	parser.AddArg("dest_files").Narg("1")
	parser.AddArg("src_files").Narg("+")

	parser.AddEnv("ENV_REQUIRED").Required()
	parser.AddEnv("ENV_OPTIONAL").Default("optional")

	config := parser.Parse()

	fmt.Fprintf(os.Stderr, "%v\n", config)
	v, ok := config.AsList("libs")
	fmt.Fprintf(os.Stderr, "%v %v\n", ok, v)
	v1, ok1 := config.AsString("input")
	fmt.Fprintf(os.Stderr, "%v %v\n", ok1, v1)
	v2, ok2 := config.AsString("output")
	fmt.Fprintf(os.Stderr, "%v %v\n", ok2, v2)
}
```
