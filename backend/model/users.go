package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
}

func (s *Store) CreateUser(user *User) error {
	return s.db.Create(user).Error
}

func (s *Store) GetUser(id uint) (*User, error) {
	var user User
	err := s.db.First(&user, id).Error
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

func (s *Store) UpdateUser(user *User) error {
	return s.db.Save(user).Error
}

func (s *Store) DeleteUser(id uint) error {
	return s.db.Delete(&User{}, id).Error
}
