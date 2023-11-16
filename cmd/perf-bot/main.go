package main

import "github.com/vvjke314/MPPR/lab1/internal/bot"

func main() {
	t := bot.NewTelegramBot()
	t.StartBot()
}
