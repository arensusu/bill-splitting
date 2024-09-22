package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	LineId   string
	Name     string
	Currency string `gorm:"default:TWD"`
	Users    []User `gorm:"many2many:members;"`
}

func (s *Store) CreateGroup(group *Group) error {
	return s.db.Create(group).Error
}

func (s *Store) GetGroup(id uint) (*Group, error) {
	var group Group
	err := s.db.First(&group, id).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (s *Store) GetGroupByLineID(lineID string) (*Group, error) {
	var group Group
	err := s.db.Where("line_id = ?", lineID).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (s *Store) ListGroupsByUserID(userID uint) ([]Group, error) {
	var groups []Group
	err := s.db.Joins("JOIN members ON members.group_id = groups.id").
		Where("members.user_id = ?", userID).
		Select("groups.id, groups.name").
		Find(&groups).Error
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (s *Store) UpdateGroup(group *Group) error {
	return s.db.Save(group).Error
}

func (s *Store) DeleteGroup(id uint) error {
	return s.db.Delete(&Group{}, id).Error
}

type Member struct {
	gorm.Model
	User    User
	UserID  uint
	Group   Group
	GroupID uint
}

func (s *Store) CreateMember(member *Member) error {
	return s.db.Create(member).Error
}

func (s *Store) GetMembership(groupID, userID uint) (*Member, error) {
	var member Member
	err := s.db.Where("group_id = ? AND user_id = ?", groupID, userID).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (s *Store) GetMember(id uint) (*Member, error) {
	var member Member
	err := s.db.First(&member, id).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (s *Store) ListMembersOfGroup(groupID uint) ([]Member, error) {
	var members []Member
	err := s.db.Joins("User").Where("group_id = ?", groupID).Find(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (s *Store) DeleteMember(groupID, userID uint) error {
	return s.db.Where("group_id = ? AND user_id = ?", groupID, userID).Delete(&Member{}).Error
}
