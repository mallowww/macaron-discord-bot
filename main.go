package main

import (
	"fmt"

	"github.com/mallowww/macaron-bot/bot"
	"github.com/mallowww/macaron-bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
