package model

import (
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var testStore *Store
var mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	db, _mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	mock = _mock

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	testStore = NewStore(gormDB)

	// testStore.db.Session(&gorm.Session{DryRun: false}).AutoMigrate(&User{}, &Group{}, &Member{}, &Expense{}, &Settlement{})

	os.Exit(m.Run())
}

func TestGormSql(t *testing.T) {
	// Load the .env file
	_ = godotenv.Load("../.env")

	store := NewStore(InitGorm())

	groups, err := store.ListUsersOfGroup(1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", groups)
}
