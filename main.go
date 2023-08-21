package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aqyuki/ssm/commands"
	"github.com/aqyuki/ssm/config"
)

func main() {

	// アプリケーションの初期化
	_, err := config.Initialize()
	if err != nil {
		if errors.Is(err, config.ErrUnsupportedOS) {
			fmt.Printf("This application supports Windows only\n")
			os.Exit(-1)
		}
		os.Exit(-1)
	}

	if err := commands.RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
