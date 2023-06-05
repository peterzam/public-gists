package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Lambda Main func
func main() {
	lambda.Start(func(ctx context.Context, snsEvent events.SNSEvent) {
		var msg string
		for _, record := range snsEvent.Records {
			msg = msg + record.SNS.Message
		}

		bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
		if err != nil {
			log.Panic(err)
		}

		chat_id, err := strconv.Atoi(os.Getenv("CHAT_ID"))
		if err != nil {
			log.Panic(err)
		}

		msg_cfg := tgbotapi.NewMessage(int64(chat_id), msg)
		bot.Send(msg_cfg)
	})
}
