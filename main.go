package main
import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"crypto-api/pkg/bitcoin"
	"crypto-api/pkg/command"
)



func NewBitcoinService() *bitcoin.BitcoinService {
	return bitcoin.NewBitcoinService()
}

func main()  {

	
	//router := gin.Default()
	bitcoinService := bitcoin.NewBitcoinService()
	commandHandler := command.NewCommandHandler(bitcoinService)

	for {
		fmt.Print("Enter the command ")
		var command string
		fmt.Scanln(&command)

		switch command {
		case "/rate":
			commandHandler.HandleRateCommand()
		default:
			fmt.Println("Unknown Command")
		}
	}

}
