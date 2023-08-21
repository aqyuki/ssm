package main

import "github.com/aqyuki/ssm/commands"

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		return
	}
}
