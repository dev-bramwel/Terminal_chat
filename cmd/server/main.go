package main

import (
	"fmt"
	"tcpchat/internal/chat"
	"tcpchat/pkg/config"
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
