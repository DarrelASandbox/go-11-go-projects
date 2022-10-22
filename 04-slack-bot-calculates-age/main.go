package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannels <-chan *slacker.CommandEvent) {
	for event := range analyticsChannels {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_SOCKET_TOKEN"))
	go printCommandEvents(bot.CommandEvents())
}
