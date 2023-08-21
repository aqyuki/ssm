package main

import (
	"os"

	"github.com/aqyuki/ssm/commands"
	"github.com/aqyuki/ssm/config"
)

func main() {

	// アプリケーションの初期化
	_, err := config.Initialize()
	if err != nil {
		os.Exit(-1)
	}

	if err := commands.RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
