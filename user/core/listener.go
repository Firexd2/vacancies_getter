package core

import (
	"context"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/uploader"
	"github.com/gotd/td/tg"
	"log"
	"math/rand"
	"os"
	"strconv"
	"vacancies_getter/structs"
)

var cvFilePath = os.Getenv("CV_PATH")

const (
	cvFileName = "beloglazov_cv.pdf"
	cvMessage  = "" +
		"Доброго времени суток!\n" +
		"Меня заинтересовала ваша вакансия Python developer.\n" +
		"Я отправляю вам свое резюме во вложении. Я уверен, что мои навыки и опыт " +
		"могут быть ценным вкладом для вашей команды. Буду рад ответить на любые вопросы или " +
		"участвовать в собеседовании."
)

func NewContactApplyListener(contacts_ch chan structs.Contact, client *telegram.Client) {
	api := tg.NewClient(client)

	for {
		select {
		case contact := <-contacts_ch:
			var userId int64
			var accessHash int64
			if contact.Username != "" {
				peer, err := api.ContactsResolveUsername(context.Background(), contact.Username)
				if err != nil {
					log.Printf("Failed to resolve a username: %v", err)
				}
				user, _ := peer.GetUsers()[0].AsNotEmpty()
				userId = user.ID
				accessHash = user.AccessHash
			} else {
				var err error
				userId, err = strconv.ParseInt(contact.Sender_id, 10, 64)
				if err != nil {
					log.Printf("Failed to parse int: %v", err)
				}
				// TODO
				accessHash = 0
			}

			// Create an instance of InputFile from the opened file
			u := uploader.NewUploader(api)
			f, err := u.FromPath(context.Background(), cvFilePath)

			if err != nil {
				log.Printf("Failed to upload file: %v", err)
			}

			// Prepare the media input for sending the file
			mediaInput := &tg.InputMediaUploadedDocument{
				File:     &tg.InputFile{ID: f.GetID(), Parts: f.GetParts()},
				MimeType: "application/pdf",
				Attributes: []tg.DocumentAttributeClass{
					&tg.DocumentAttributeFilename{
						FileName: cvFileName,
					},
				},
			}

			_, err = api.MessagesSendMedia(context.Background(), &tg.MessagesSendMediaRequest{
				Peer:     &tg.InputPeerUser{UserID: userId, AccessHash: accessHash},
				Message:  cvMessage,
				RandomID: rand.Int63(),
				Media:    mediaInput,
			})
			if err != nil {
				log.Printf("Failed to send message: %v", err)
			}
		}
	}
}
