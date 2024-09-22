package model

import "gorm.io/gorm"

type GroupInvitation struct {
	gorm.Model
	Code    string
	GroupID uint
	Group   Group
}

func (s *Store) CreateGroupInvitation(invitation *GroupInvitation) error {
	return s.db.Create(invitation).Error
}

func (s *Store) GetGroupInvitation(code string) (*GroupInvitation, error) {
	var invitation GroupInvitation
	err := s.db.Where("code = ?", code).First(&invitation).Error
	if err != nil {
		return nil, err
	}
	return &invitation, nil
}

func (s *Store) DeleteGroupInvitation(code string) error {
	return s.db.Where("code = ?", code).Delete(&GroupInvitation{}).Error
}
