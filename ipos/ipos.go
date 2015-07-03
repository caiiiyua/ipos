package ipos

import (
	"errors"
	"time"
)

type Pos struct {
	StaffId string `json:"staffId"`
	UserId  string `json:"userId"`
	Type    int    `json:"type"`
	Items   []struct {
		ItemId string `json:"itemId"`
		Qty    int    `json:"qty"`
	} `json:"items"`
}

func NewPos(staffId string) (*Pos, error) {
	if staffId == "" || len(staffId) != 2 {
		return nil, errors.New("empty or invalid staff Id")
	}
	return &Pos{StaffId: staffId}, nil
}

type PosManager struct {
}

func NewPosManager() *PosManager {
	return nil
}

func (m *PosManager) Save(pos *Pos) error {
	return nil
}

func (m *PosManager) Delete(pos *Pos) error {
	return nil
}

func (m *PosManager) List() ([]*Pos, error) {
	return nil, nil
}

type UserItem struct {
	Id        int64
	UserId    int64  // related user id
	CardId    string // related user's card no
	ItemId    int64  // related item id
	Qty       int64
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

type Item struct {
	Id          int64
	Code        string `xorm:"unique not null index"`
	Name        string
	Description string `xorm:"varchar(64)"`
	Price       float64
}
