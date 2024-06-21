package rest

import (
	"fmt"
	"net/http"
	"validator/records"

	"github.com/gin-gonic/gin"
)

type userData struct {
	CardNumber string `json:"cardNumber" binding:"required"`
}

type cardMessage struct {
	Brand   string `json:"brand"`
	Type    string `json:"type"`
	Country string `json:"country"`
}

func Run() {
	router := gin.Default()

	router.POST("/validate", postValidate)

	router.Run("localhost:8080")
}

func postValidate(c *gin.Context) {
	var newUserData userData

	if err := c.BindJSON(&newUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if !isValidLuhn(newUserData.CardNumber) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid card number"})

		return
	}

	record, err := records.GetRecords()
	if err != nil {
		return
	}

	data := records.GetBinDataFromRecord(newUserData.CardNumber, record)

	var receivedData cardMessage
	if data != nil {
		receivedData = cardMessage{
			Brand:   data[records.Brand],
			Type:    data[records.Type],
			Country: data[records.Country]}
	} else {
		receivedData = cardMessage{
			Brand: fmt.Sprint(getMII(newUserData.CardNumber))}
	}

	c.JSON(http.StatusOK, receivedData)
}

func isValidLuhn(cardNumber string) bool {
	length := len(cardNumber)
	total := 0
	even := true

	for i := length - 2; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if digit < 0 || digit > 9 {
			return false
		}

		if even {
			digit <<= 1
		}

		even = !even

		if digit > 9 {
			total += digit - 9
		} else {
			total += digit
		}
	}
	checkSum := int(cardNumber[length-1] - '0')

	return (total+checkSum)%10 == 0
}

type MII int

const (
	AmericanExpress = iota + 3
	Visa
	MasterCard
	Discover
)

var MIIName = map[MII]string{
	AmericanExpress: "American Express",
	Visa:            "Visa",
	MasterCard:      "MasterCard",
	Discover:        "Discover",
}

func (mii MII) String() string {
	return MIIName[mii]
}

func getMII(cardNumber string) MII {
	return MII(int(cardNumber[0] - '0'))
}
