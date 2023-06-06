package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"vacancies_getter/bot"
	"vacancies_getter/structs"
	"vacancies_getter/user"
)

func main() {
	new_vanacies_ch := make(chan structs.NewVacancy)
	contacts_ch := make(chan structs.Contact)
	go user.Main(new_vanacies_ch, contacts_ch)
	go bot.Main(new_vanacies_ch, contacts_ch)

	// Will block here until user hits ctrl+c
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Blocking, press ctrl+c to continue...")
	<-done
}
