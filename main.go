package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/kevinliu399/orders-api/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel() // cancel when we are finished

	err := app.Start(ctx)
	if err != nil {
		fmt.Println(err)
	}

}
