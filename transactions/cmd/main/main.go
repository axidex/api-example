package main

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/transactions/internal/app"
)

func main() {
	ctx := context.Background()

	if err := app.NewApp().Run(ctx); err != nil {
		fmt.Println(err)
		return
	}
}
