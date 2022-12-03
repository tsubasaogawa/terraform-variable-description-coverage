// main package is the main of tfmodblock.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

// Variable is Terraform variable object.
type Variable struct {
	Name        string
	Description string
}

const VERSION string = "0.0.0"

func main() {
	var (
		v = flag.Bool("v", false, "tfvdc version")
	)
	flag.Parse()

	if *v {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	path := "."
	if flag.NArg() > 0 {
		path = flag.Arg(0)
	}

	if !tfconfig.IsModuleDir(path) {
		panic("given path does not contain tf files")
	}

	module, _ := tfconfig.LoadModule(path)
	for k, v := range module.Variables {
		hasDescription := v.Description != ""
		fmt.Printf("%s: %t\n", k, hasDescription)
	}
	// fmt.Println(block)
}
