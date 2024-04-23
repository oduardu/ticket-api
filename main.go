package main

import (
	"fmt"

	"github.com/oduardu/ticket-api/config"
	"github.com/oduardu/ticket-api/router"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Errorf("Config Initialization error: %v", err)
		return
	}

	router.Initialize()
}
