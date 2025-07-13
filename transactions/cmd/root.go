package cmd

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/transactions/internal/app"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "TonApp",
	Short: "Wrapper to run TonTransactionsListener or TonAPI",
}

func init() {
	RootCmd.AddCommand(transactionsCmd)
	RootCmd.AddCommand(apiCmd)
}

var transactionsCmd = &cobra.Command{
	Use:   "transactions",
	Short: "Run TonTransactionsListener",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		if err := app.NewTransactionsApp().Run(ctx); err != nil {
			fmt.Println("transactions app error:", err)
			return
		}
	},
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run ApiApp",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		if err := app.NewApiApp().Run(ctx); err != nil {
			fmt.Println("api app error:", err)
			return
		}
	},
}
