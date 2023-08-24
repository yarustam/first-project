package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

var saveMap = make(map[int64]string)

//в мапе int64 тип ключа, string тип значения, ключом у нас будет уникальный id пользователя

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
	})

	switch {
	case strings.HasPrefix(update.Message.Text, "save"): //проверяет начало строки с указанного префикса
		chatID := update.Message.Chat.ID                               //тут мы присваим значение id пользователя
		saveMesUser := strings.TrimPrefix(update.Message.Text, "save") //удаляет указанный префикс из переданной строки (если он есть)
		saveMap[chatID] = saveMesUser                                  // сохраняем текст для текущего пользователя

	case strings.HasPrefix(update.Message.Text, "get"):
		chatID := update.Message.Chat.ID
		saveMesUser := "No saved words"

		if savedText, ok := saveMap[chatID]; ok { //проверка на наличие сохраненного текста пользователя
			saveMesUser = savedText
		}

		b.SendMessage(ctx, &bot.SendMessageParams{ //вывод, если сохраненного текста не было, выведем присвоенное ранее "No saved words"
			ChatID: update.Message.Chat.ID,
			Text:   saveMesUser,
		})
	}
}
