package main

import (
	"helmvmgr/pkg/cmdcli"
	"os"
)

func main() {
	cmd := cmdcli.NewCliCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
