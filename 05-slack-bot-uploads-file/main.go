package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN_2"))
	channelArr := []string{os.Getenv("SLACK_CHANNEL_ID")}
	fileArr := []string{"05-slack-bot-uploads-file/rubber-duck-g2e08fca78_640.jpg"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
