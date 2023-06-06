package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"regexp"
	"vacancies_getter/bot/core"
	"vacancies_getter/structs"
)

func Main(new_vanacies_ch chan structs.NewVacancy, contacts_ch chan structs.Contact) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	go core.NewVacanciesChListener(new_vanacies_ch, bot)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			r, _ := regexp.Compile("@[\\w]+")
			username := r.FindString(update.CallbackQuery.Message.Text)

			if username != "" {
				// username looks like @firexd2 and we don't need @
				username = username[1:]
			}

			contacts_ch <- structs.Contact{username, update.CallbackQuery.Data}

			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "ok")
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}
		}
	}
}
