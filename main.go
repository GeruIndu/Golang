package main

import (
	"context"
	"fmt"

	application "github.com/GeruIndu/Golang_routing_module/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start error : ", err)
	}
}
