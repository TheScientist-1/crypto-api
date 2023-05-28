package routes
import ("net/http"
		"github.com/gin-gonic/gin"
		"net/url"
		"io/ioutil"
		"encoding/json"
)

const(
	baseUrl = "https://api.binance.com"
	pathEndpoint = "/api/v3/avgPrice"
)
func GetPrice(c *gin.Context) {
	_url, _ := url.ParseRequestURI(baseUrl)
	_url.Path = pathEndpoint
	parameters := url.Values{}
	parameters.Add("symbol", "BTCUAH")
	_url.RawQuery = parameters.Encode()

	response, _ := http.Get(_url.String())
	responseBody, _ := ioutil.ReadAll(response.Body)

	price := struct {
		Minutes int    `json:"mins"`
		Price   string `json:"price"`
	}{}
	json.Unmarshal(responseBody, &price)

	c.JSON(http.StatusOK, gin.H{"rate": price.Price})
}
