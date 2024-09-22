package model

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	MemberID         uint
	Member           Member
	OriginalAmount   float64
	OriginalCurrency string `gorm:"default:TWD"`
	ConvertedAmount  float64
	Category         string
	Description      string
	Date             time.Time
	IsSettled        bool `gorm:"default:false"`
}

func (s *Store) CreateExpense(expense *Expense) error {
	return s.db.Create(expense).Error
}

func (s *Store) GetExpense(id uint) (*Expense, error) {
	var expense Expense
	err := s.db.First(&expense, id).Error
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func (s *Store) ListExpenses(groupID uint) ([]Expense, error) {
	var expenses []Expense
	err := s.db.Joins("Member").Where("Member.group_id = ?", groupID).Find(&expenses).Error
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func (s *Store) ListNonSettledExpenses(groupID uint) ([]Expense, error) {
	var expenses []Expense
	err := s.db.Joins("Member").Where("Member.group_id = ? AND is_settled = ?", groupID, false).Find(&expenses).Error
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func (s *Store) UpdateExpense(expense *Expense) error {
	return s.db.Save(expense).Error
}

func (s *Store) DeleteExpense(id uint) error {
	return s.db.Delete(&Expense{}, id).Error
}

func (s *Store) ListExpensesWithinDate(groupID uint, startTime, endTime time.Time) ([]Expense, error) {
	var expenses []Expense
	err := s.db.Joins("Member").Where("Member.group_id = ? AND date BETWEEN ? AND ?", groupID, startTime, endTime).Find(&expenses).Error
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

type ExpenseSummary struct {
	Category string
	Total    float64
}

func (s *Store) SummarizeExpensesWithinDate(groupID uint, startTime, endTime time.Time) ([]ExpenseSummary, error) {
	var summaries []ExpenseSummary
	err := s.db.Table("expenses").
		Select("category, SUM(converted_amount) as total").
		Joins("JOIN members ON expenses.member_id = members.id").
		Where("members.group_id = ? AND expenses.date BETWEEN ? AND ?", groupID, startTime, endTime).
		Group("category").
		Scan(&summaries).Error
	if err != nil {
		return nil, err
	}
	return summaries, nil
}
