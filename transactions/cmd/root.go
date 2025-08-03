package cmd

import (
	"context"
	"fmt"
	fxApp "github.com/axidex/api-example/transactions/internal/fx"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
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
		app := fxApp.NewTransactionsApp()
		ctx := context.Background()

		if err := app.Start(ctx); err != nil {
			fmt.Println("transactions app start error:", err)
			return
		}

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		if err := app.Stop(ctx); err != nil {
			fmt.Println("transactions app stop error:", err)
		}
	},
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run ApiApp",
	Run: func(cmd *cobra.Command, args []string) {
		app := fxApp.NewApiApp()
		ctx := context.Background()

		if err := app.Start(ctx); err != nil {
			fmt.Println("api app start error:", err)
			return
		}

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig

		if err := app.Stop(ctx); err != nil {
			fmt.Println("api app stop error:", err)
		}
	},
}
