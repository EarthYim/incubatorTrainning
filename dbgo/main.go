package main

import (
	"database/sql"
	"dbgo/wallet"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	//wallet.Conn, err = sql.Open("postgres", "postgres://ivzeojed:0IZOHrr423bv7Mf-W9cGGozq1NPfJYBf@tiny.db.elephantsql.com/ivzeojed")
	wallet.Conn, err = sql.Open("postgres", "postgres://earth:12321@0.0.0.0:5432/db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	//defer wallet.Conn.Close()

	if err = wallet.Conn.Ping(); err != nil {
		log.Fatal(err)
	}

	sql := `CREATE TABLE IF NOT EXISTS wallet(
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			balance DECIMAL(10,2) NOT NULL DEFAULT 0.00
		)`

	_, err = wallet.Conn.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	r := newServer()

	// shutdown := make(chan os.Signal, 1)
	// signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// go func() {

	// }

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
