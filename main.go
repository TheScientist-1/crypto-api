package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"crypto-api/pkg/bitcoin"
	"crypto-api/pkg/command"
)



func NewBitcoinService() *bitcoin.BitcoinService {
	return bitcoin.NewBitcoinService()
}

func main()  {

	
	router := gin.Default()
	bitcoinService := NewBitcoinService()
	commandHandler := NewCommandHandler(bitcoinService)

	for {
		fmt.Print("Введіть команду: ")
		var command string
		fmt.Scanln(&command)

		switch command {
		case "/rate":
			commandHandler.HandleRateCommand()
		default:
			fmt.Println("Невідома команда")
		}
	}

}
