package model

import (
	"testing"

	"gorm.io/gorm"
)

func TestDB(t *testing.T) {
	store := NewStore()
	store.db.Session(&gorm.Session{DryRun: true}).AutoMigrate(&User{}, &Group{}, &Member{}, &Expense{}, &Settlement{})
}
