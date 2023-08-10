package wallet

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Wallet struct {
	ID      string  `json:"id"`
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

type Request struct {
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

type amount struct {
	Amount float64
}

var wallets = make(map[string]Wallet) //for storing wallet in memory

func CreateWalletHandler(c *gin.Context) {
	wt := &Wallet{}
	//request := &Request{}

	err := c.ShouldBindJSON(wt)
	if err != nil {
		log.Println("err")
	}

	wt.ID = uuid.NewString()

	wallets[wt.ID] = *wt
	log.Printf("%#v\n", *wt)
	c.JSON(200, *wt)
}

func GetWalletByIDHandler(c *gin.Context) {
	id := c.Param("id")
	wt, ok := wallets[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "wallet not found!",
		})
	}
	c.JSON(200, wt)
}

func GetBalanceByIDHandler(c *gin.Context) {
	id := c.Param("id")
	wt, ok := wallets[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "wallet not found!",
		})
	}
	c.JSON(200, gin.H{
		"id":      wt.ID,
		"balance": wt.Balance,
	})
}

func DepositByIDHandler(c *gin.Context) {
	id := c.Param("id")
	wt, ok := wallets[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "wallet not found!",
		})
	}

	// dep := &amount{}
	var dep struct {
		Amount float64 `json:"amount"`
	}
	err := c.ShouldBindJSON(&dep)
	if err != nil {
		log.Println("err")
	}
	log.Println("Deposit Amount:", dep.Amount)

	wt.Balance = wt.Balance + dep.Amount
	wallets[wt.ID] = wt
	log.Println("Balance:", wt.Balance)
	c.JSON(200, gin.H{
		"id":      wt.ID,
		"balance": wt.Balance,
	})
}

func WithdrawByIDHandler(c *gin.Context) {
	id := c.Param("id")
	wt, ok := wallets[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "wallet not found!",
		})
	}

	// wit := &Amount{}
	var wit struct {
		Amount float64 `json:"amount"`
	}
	err := c.ShouldBindJSON(&wit)
	if err != nil {
		log.Println("err")
	}
	log.Println("Withdraw Amount:", wit.Amount)

	if wt.Balance = wt.Balance - wit.Amount; wt.Balance < 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Insuffucient Fund",
		})
		wt.Balance = wt.Balance + wit.Amount
	}
	log.Println("Balance:", wt.Balance)
	wallets[wt.ID] = wt
	c.JSON(200, gin.H{
		"id":      wt.ID,
		"balance": wt.Balance,
	})
}
