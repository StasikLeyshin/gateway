package main

import (
	"context"
	"fmt"
	"gateway/internal/app"
)

func main() {
	startApp, err := app.NewApp(context.Background())
	if err != nil {
		_ = fmt.Errorf("failed to app: %v", err)
		return
	}

	startApp.Run(context.Background())
}
