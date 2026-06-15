package main

import (
	"fmt"
	"net/internal/chat"
	"net/pkg/config"
)

func main() {
	cfg, ok := config.MustParse()
	if !ok {
		return
	}

	if err := chat.ValidatePort(cfg.Port); err != nil {
		fmt.Println(config.Usage)
		return
	}

	server := chat.NewServer(cfg.Port)
	if err := server.Start(); err != nil {
		fmt.Println(err)
	}
}
