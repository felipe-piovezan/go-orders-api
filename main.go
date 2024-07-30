package main

import (
	"context"
	"fmt"

	"github.com/felipe-piovezan/go-orders-api/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("falied to start app:", err)
	}
}
