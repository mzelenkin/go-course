package main

import (
	"fmt"
	"os"
)

var envDir string
var program string
var programArgs []string

func init() {
	args := os.Args
	if len(args) < 3 {
		println("Usage: ", args[0], " path program")
	}

	envDir = args[1]
	program = args[2]
	programArgs = args[3:]
	println(program)
}

func main() {
	if _, err := os.Stat(envDir); os.IsNotExist(err) {
		fmt.Printf("Path %s does not exists\n", envDir)
		os.Exit(2)
	}

	vars, err := loadEnv(envDir)
	if err != nil {
		fmt.Println("Loading variables error: ", envDir)
		os.Exit(2)
	}

	runInEnv(program, programArgs, mergeEnv(vars))
}
