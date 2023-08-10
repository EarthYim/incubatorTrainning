package main

import (
	"encoding/json"
	"log"
)

// type wallet struct {
// 	ID      string  `json:"id"`
// 	Owner   string  `json:"owner"`
// 	Balance float64 `json:"balance"`
// }

func main() {
	b := `{"ID": "wallet_Earth", "Owner": "Earth", "Balance": 100.0}`
	wt := &Wallet{}
	err := json.Unmarshal([]byte(b), wt)
	if err != nil {
		log.Println("err:", err)
		return
	}

	log.Printf("wt: %#v\n", wt)

	out, err := json.Marshal(wt)
	if err != nil {
		log.Println("Err")
		return
	}
	log.Println("out:", string(out))

}
