package wallet

import (
	"database/sql/driver"
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
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println("error creating mock database")
		return
	}
	defer db.Close()

	query := "INSERT INTO wallet"
	expectedArgs := []driver.Value{"", 0.0}

	// Expect QueryRow to be called with the provided query and args
	mock.ExpectQuery(query).WithArgs(expectedArgs...).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(-1)) // Mock returned row

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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("error creating mock databse")
	}

	defer db.Close()

	query := "INSERT INTO wallet"
	var Owner = StubOwner{
		owner:   "Earth",
		balance: 100.0,
	}
	expectedArgs := []driver.Value{Owner.owner, Owner.balance}

	// Expect QueryRow to be called with the provided query and args
	mock.ExpectQuery(query).WithArgs(expectedArgs...).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1)) // Mock returned row

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

func TestInsertWalletFake(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("error creating mock database")
	}
	defer db.close()

	query := "INSERT INTO wallet"
}
