package command

import (
	"fmt"
	"crypto-api/pkg/bitcoin"
)

type CommandHandler struct {
	bitcoinService *bitcoin.BitcoinService
}

func NewCommandHandler(bitcoinService *bitcoin.BitcoinService) *CommandHandler {
	return &CommandHandler{
		bitcoinService: bitcoinService,
	}
}

func (ch *CommandHandler) HandleRateCommand() {
	price, err := ch.bitcoinService.GetCurrentPrice()
	if err != nil {
		fmt.Println("Помилка при отриманні курсу біткоїна:", err)
		return
	}

	fmt.Printf("Поточний курс біткоїна до UAH: %.2f\n", price)
}
