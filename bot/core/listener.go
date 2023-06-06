package core

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"vacancies_getter/structs"
)

const myUserId = 661703753

func NewVacanciesChListener(new_vanacies_ch chan structs.NewVacancy, bot *tgbotapi.BotAPI) {
	for {
		select {
		case vacancy := <-new_vanacies_ch:
			msg := tgbotapi.NewMessage(myUserId, vacancy.Text)
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Откликнуться", string(rune(vacancy.Sender_id))),
				),
			)
			bot.Send(msg)
		}
	}
}
