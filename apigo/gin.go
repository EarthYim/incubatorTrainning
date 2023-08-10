package main

import (
	"apigo/wallet"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := newServer()
	r.Run()

}

// Logger handler Middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("handler:", c.Request.Method, c.Request.URL)
		c.Next()
		log.Println("after call next handler:", c.Request.Method, c.Request.URL)
	}
}

func newServer() *gin.Engine {

	r := gin.Default()

	r.POST("/wallet", wallet.CreateWalletHandler)
	r.GET("/wallets/:id", wallet.GetWalletByIDHandler)
	r.GET("/wallets/:id/balance", wallet.GetBalanceByIDHandler)
	r.POST("/wallets/:id/deposit", wallet.DepositByIDHandler)
	r.POST("/wallets/:id/withdraw", wallet.WithdrawByIDHandler)

	return r
}
