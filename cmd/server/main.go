package main

import (
	"fmt"

	"github.com/Pholluxion/task-api/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		fmt.Println("❌ Failed to initialize app:", err)
		return
	}

	if err := a.Start(); err != nil {
		fmt.Println("❌ Failed to start server:", err)
	}
}
