package telegram

import "github.com/15683/bot/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func new(client *telegram.Client) {

}
