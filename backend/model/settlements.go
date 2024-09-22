package model

import "gorm.io/gorm"

type Settlement struct {
	gorm.Model
	ID      uint `gorm:"autoIncrement"`
	Payer   Member
	PayerID uint `gorm:"primaryKey"`
	Payee   Member
	PayeeID uint `gorm:"primaryKey"`
	Amount  float64
}

func (s *Store) CreateSettlement(settlement *Settlement) error {
	return s.db.Create(settlement).Error
}

func (s *Store) GetSettlement(payerID, payeeID uint) (*Settlement, error) {
	var settlement Settlement
	err := s.db.Where("payer_id = ? AND payee_id = ?", payerID, payeeID).First(&settlement).Error
	if err != nil {
		return nil, err
	}
	return &settlement, nil
}

func (s *Store) ListSettlements(groupID uint) ([]Settlement, error) {
	var settlements []Settlement
	err := s.db.Joins("Payer").Joins("Payee").
		Where("Payer.group_id = ? OR Payee.group_id = ?", groupID, groupID).
		Select("settlements.payer_id, settlements.payee_id, settlements.amount").
		Find(&settlements).Error
	if err != nil {
		return nil, err
	}
	return settlements, nil
}

func (s *Store) UpdateSettlement(settlement *Settlement) error {
	return s.db.Where("payer_id = ? AND payee_id = ?", settlement.PayerID, settlement.PayeeID).
		Updates(Settlement{Amount: settlement.Amount}).Error
}

func (s *Store) DeleteSettlement(payerID, payeeID uint) error {
	return s.db.Where("payer_id = ? AND payee_id = ?", payerID, payeeID).Delete(&Settlement{}).Error
}
