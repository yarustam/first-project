package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {

	ctx := context.Background()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New("6689732780:AAFWWLcbl91Xhh_K3_-MDDJzzDKGqR68Hvs", opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.Text == "/start" {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Greetings! I'm a little wolf bot, nice to meet you!I have know two command:\n\n/save + you'r message - I will remember it\n/get - I will write the message, which I have remember",
		})
	} else {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   update.Message.Text,
		})
	}
	fmt.Println(update.Message.Text)
}
