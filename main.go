package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

var mytext = "empty"

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
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		//Text: update.Message.Text,
	})

	switch {
	case strings.HasPrefix(update.Message.Text, "save"):
		mytext = strings.TrimPrefix(update.Message.Text, "save")
	case strings.HasPrefix(update.Message.Text, "get"):
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   mytext,
		})
	}

}
