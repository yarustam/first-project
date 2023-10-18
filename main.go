package main

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

var saveMap = make(map[int64][]string)

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

	switch {
	case strings.HasPrefix(update.Message.Text, "save"): //проверяет начало строки с указанного префикса
		chatID := update.Message.Chat.ID                                //тут мы присваим значение id пользователя
		saveMesUser := strings.TrimPrefix(update.Message.Text, "save ") //удаляет указанный префикс из переданной строки (если он есть)
		saveMap[chatID] = append(saveMap[chatID], saveMesUser)          //тут мы при помощи функции append добавляем значение в слайс                                  //сохраняем текст для текущего пользователя

	case strings.HasPrefix(update.Message.Text, "get"):
		chatID := update.Message.Chat.ID
		saveMesUser := "No saved words"

		if savedText, ok := saveMap[chatID]; ok { //проверка на наличие сохраненного текста пользователя
			saveMesUser = savedText[len(savedText)-1] //такая конструкция позволяет выводить только последний элемент слайса
		}

		b.SendMessage(ctx, &bot.SendMessageParams{ //вывод, если сохраненного текста не было, выведем присвоенное ранее "No saved words"
			ChatID: update.Message.Chat.ID,
			Text:   saveMesUser,
		})

	case strings.HasPrefix(update.Message.Text, "all"):
		chatID := update.Message.Chat.ID
		saveMesUser := "No saved words"

		if savedText, ok := saveMap[chatID]; ok {
			saveMesUser = strings.Join(savedText, "\n")
		}
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   saveMesUser,
		})
	default:
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Будьте здоровы! Please enter correct command:\n\nsave, чтобы сохранить сообщение\nget, вывести последнее сообщение \nall, вывести все сообщения\n\n(я еще маленький бот, но уже не пугаюсь сов🦉)",
		})
	}
}
