package wallet

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// func (db *DummyDB) QueryRow(query string, owner string, balance float64) DummyDB {
// 	return DummyDB{}
// }

// type DBer interface {
// 	QueryRow(query string, owner string, balance float64) DummyDB
// }

func TestInserWalletDummy(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}
	var id int
	err, id = InsertWallet(db, "", 0)
	if err != nil {
		t.Error(err)
	}
	want := -1
	if id != want {
		t.Errorf("Want '%d', got '%d'", want, id)
	}
}

type StubOwner struct {
	owner   string
	balance float64
}

func TestInsertWalletStub(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("error creating mock databse")
	}

	var Owner = StubOwner{
		owner:   "Earth",
		balance: 100.0,
	}

	var id int
	err, id = InsertWallet(db, Owner.owner, Owner.balance)
	if err != nil {
		t.Error(err)
	}

	want := 1
	if id != want {
		t.Errorf("Want '%d', got '%d'", want, id)
	}
}
