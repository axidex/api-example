package main

import (
	"github.com/axidex/api-example/transactions/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		return
	}
}
