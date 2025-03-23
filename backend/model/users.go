package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	LineID   string `gorm:"uniqueIndex:composite_index"`
	Username string

	DiscordId string `gorm:"uniqueIndex:composite_index"`
}

func (s *Store) CreateUser(user *User) error {
	return s.db.Create(user).Error
}

func (s *Store) GetUserByLineID(lineId string) (*User, error) {
	var user User
	err := s.db.Where("line_id = ?", lineId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Store) GetUserByUsername(username string) (*User, error) {
	var user User
	err := s.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Store) DeleteUser(id uint) error {
	return s.db.Delete(&User{}, id).Error
}

func (s *Store) GetUserByDiscordID(discordId string) (*User, error) {
	var user User
	err := s.db.Where("discord_id = ?", discordId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
