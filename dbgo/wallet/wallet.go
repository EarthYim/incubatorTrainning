package wallet

import (
	"database/sql"
	"dbgo/db"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

// var wallets = make(map[string]Wallet) //for storing wallet in memory
var Conn *sql.DB

func CreateWalletHandler(c *gin.Context) {

	request := &Request{}
	wt := Wallet{}

	err := c.ShouldBindJSON(request)
	if err != nil {
		log.Fatal(err)
	}

	// wt.ID = uuid.NewString()

	var id int
	err, id = db.InsertWallet(Conn, request.Owner, request.Balance)
	if err != nil {
		log.Fatal(err)
	}

	wt.Owner = request.Owner
	wt.Balance = request.Balance
	wt.ID = strconv.Itoa(id)

	log.Printf("%#v\n", wt)
	c.JSON(200, wt)
}

func GetWalletByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	wt := db.Wallet{}
	// wt, ok := wallets[id]
	// if !ok {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error": "wallet not found!",
	// 	})
	// }

	err, wt = db.SelectWalletById(Conn, id)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, wt)
}

func GetBalanceByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	wt := db.Wallet{}
	// wt, ok := wallets[id]
	// if !ok {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"error": "wallet not found!",
	// 	})
	// }
	err, wt = db.SelectWalletById(Conn, id)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"id":      wt.ID,
		"balance": wt.Balance,
	})
}

func DepositByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	var dep struct {
		Amount float64
	}

	err = c.ShouldBindJSON(&dep)
	if err != nil {
		log.Println("err")
	}
	log.Println("Deposit Amount:", dep.Amount)

	wt := db.Wallet{}

	err, wt = db.SelectWalletById(Conn, id)
	if err != nil {
		log.Fatal(err)
	}

	wt.Balance = wt.Balance + dep.Amount

	err = db.UpdateWalletBalance(Conn, id, wt.Balance)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Balance:", wt.Balance)
	c.JSON(200, gin.H{
		"id":      wt.ID,
		"balance": wt.Balance,
	})
}

func WithdrawByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	var wit struct {
		Amount float64
	}

	err = c.ShouldBindJSON(&wit)
	if err != nil {
		log.Println("err")
	}
	log.Println("Withdraw Amount:", wit.Amount)

	wt := db.Wallet{}

	err, wt = db.SelectWalletById(Conn, id)
	if err != nil {
		log.Fatal(err)
	}

	if wt.Balance = wt.Balance - wit.Amount; wt.Balance < 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Insuffucient Fund",
		})
		wt.Balance = wt.Balance + wit.Amount
	} else {
		err = db.UpdateWalletBalance(Conn, id, wt.Balance)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Balance:", wt.Balance)
		c.JSON(200, gin.H{
			"id":      wt.ID,
			"balance": wt.Balance,
		})
	}

	log.Println("Balance:", wt.Balance)
}
