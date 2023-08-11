package db

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type Wallet struct {
	ID      string
	Owner   string
	Balance float64
}

func InsertWallet(db *sql.DB, owner string, balance float64) (error, int) {
	query := `INSERT INTO wallet (name, balance)
			VALUES ($1, $2) RETURNING id`

	row := db.QueryRow(query, owner, balance)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("last insert id: ", id)
	return nil, id
}

func UpdateWalletBalance(db *sql.DB, id int, amount float64) error {
	query := `UPDATE wallet SET balance = $1 WHERE id = $2`
	_, err := db.Exec(query, amount, id)
	if err != nil {
		return err
	}
	return nil
}

// func selectAllWallet(db *sql.DB) error {
// 	query := `SELECT * FROM wallet`
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	var ws []Wallet

// 	for rows.Next() {
// 		var id int
// 		var owner string
// 		var balance float64
// 		err = rows.Scan(&id, &owner, &balance)
// 		if err != nil {
// 			return err
// 		}

// 		wt := Wallet{
// 			ID:      id,
// 			Owner:   owner,
// 			Balance: balance,
// 		}

// 		ws = append(ws, wt)

// 		log.Println("id:", id, "owner:", owner, "balance:", balance)
// 	}

// 	log.Printf("wallets: %#v\n", ws)
// 	return nil
// }

func SelectWalletById(db *sql.DB, walletID int) (error, Wallet) {
	wt := Wallet{}
	query := `SELECT * FROM wallet WHERE id = $1`
	row := db.QueryRow(query, walletID)
	var id int
	var owner string
	var balance float64
	err := row.Scan(&id, &owner, &balance)
	if err != nil {
		return err, wt
	}

	wt.ID = strconv.Itoa(id)
	wt.Owner = owner
	wt.Balance = balance

	log.Println("id:", id, "owner:", owner, "balance:", balance)
	return nil, wt
}

func deleteWallet(db *sql.DB, walletID int) error {
	query := `DELETE FROM wallet WHERE id = $1`
	_, err := db.Exec(query, walletID)
	if err != nil {
		return err
	}
	return nil
}

// func main() {
// 	conn, err := sql.Open("postgres", "postgres://ivzeojed:0IZOHrr423bv7Mf-W9cGGozq1NPfJYBf@tiny.db.elephantsql.com/ivzeojed")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer conn.Close()

// 	if err = conn.Ping(); err != nil {
// 		log.Fatal(err)
// 	}

// 	sql := `CREATE TABLE IF NOT EXISTS wallet(
// 			id SERIAL PRIMARY KEY,
// 			name VARCHAR(255) NOT NULL,
// 			balance DECIMAL(10,2) NOT NULL DEFAULT 0.00
// 		)`

// 	_, err = conn.Exec(sql)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// err = insertWallet(conn, "John Doe", 100.00)
// if err != nil {
// 	log.Fatal("insert error:", err)
// }

// err = updateWalletBalance(conn, 1, 200.00)
// if err != nil {
// 	log.Fatal("update error:", err)
// }

// err = selectAllWallet(conn)
// if err != nil {
// 	log.Fatal("select all error:", err)
// }

// err = selectWalletById(conn, 1)
// if err != nil {
// 	log.Fatal("select by id error:", err)
// }

// err = deleteWallet(conn, 3)
// if err != nil {
// 	log.Fatal("delete error:", err)
// }
// }
