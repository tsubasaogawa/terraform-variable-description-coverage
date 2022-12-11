// main package is the main of tfvdc
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

const VERSION string = "0.0.0"

// getCol returns column number of the error line (TODO)
func getCol(filename string, varName string) int {
	return 1
}

// errorformat returns `errorformat` style like https://github.com/reviewdog/errorformat
func errorformat(filename string, line int, col int, msg string) string {
	return fmt.Sprintf("%s:%d:%d: %s\n", filename, line, col, msg)
}

// main function
func main() {
	var (
		varMode    = flag.Bool("v", true, "variable mode")
		outputMode = flag.Bool("o", false, "output mode (TODO)")
		version    = flag.Bool("version", false, "tfvdc version")
	)
	flag.Parse()

	if *varMode && *outputMode {
		panic("Choose either `-v` (default) or `-o`")
	} else if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	path := "."
	if flag.NArg() > 0 {
		path = flag.Arg(0)
	}

	if !tfconfig.IsModuleDir(path) {
		panic("Given path does not contain tf files")
	}

	module, _ := tfconfig.LoadModule(path)
	for k, v := range module.Variables {
		hasDescription := v.Description != ""
		if hasDescription {
			msg := "variable `" + k + "` does not have description"
			print(errorformat(v.Pos.Filename, v.Pos.Line, getCol(v.Pos.Filename, k), msg))
		}
	}
}
