package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello, World!"))
// }

/*
POST /wallets
GET /wallets/:id/balance
POST /wallets/:id/deposit
POST /wallets/:id/withdraw

```json
POST /wallets
{
  "id": "wallet_AnuchitO",
  "owner": "AnuchitO",
  "balance": 100.0,
}
```
*/

// type wallet struct {
// 	ID      string
// 	Owner   string
// 	Balance float64
// }

func main() {

	log.Println("Starting Server")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

		name := r.URL.Query().Get("name")
		fmt.Println("Name is: ", name)

		if r.Method == "GET" {
			w.Write([]byte("Hello, World!"))
			return
		}

		if r.Method == "POST" {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			log.Println("body:", string(b))
			wt := Wallet{}
			if err := json.Unmarshal(b, &wt); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			log.Printf("%#v\n", wt)

			out, err := json.Marshal(wt)
			if err != nil {
				log.Println("Marshal err")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(out)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	// err := http.ListenAndServe(":12345678", nil)
	// if err != nil {
	// 	//log.Println(err)
	// 	log.Fatal("ListenAndServe: ", err)
	// }
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("Bye")
}
