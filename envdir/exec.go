package main

import (
	"log"
	"os"
	"os/exec"
)

// mergeEnv склеивает текущие переменные окружения с переменными в vars
// и их возвращает в виде []string{"VAR=VALUE"}
func mergeEnv(vars map[string]string) []string {
	var newEnv []string
	for varName, varValue := range vars {
		newEnv = append(newEnv, varName+"="+varValue)
	}
	newEnv = append(newEnv, os.Environ()...)

	return newEnv
}

// runInEnv запускает на выполенение программу в окружении переменных env
func runInEnv(program string, args []string, env []string) {
	cmd := exec.Command(program, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}
