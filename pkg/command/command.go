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
		fmt.Println("Error getting Bitcoin rate:", err)
		return
	}

	fmt.Printf("Current Bitcoin exchange rate to UAH: %.2f\n", price)
}
