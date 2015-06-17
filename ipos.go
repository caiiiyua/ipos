package ipos

import "errors"

type Pos struct {
	Id string
}

func NewPos(id string) (*Pos, error) {
	if id == "" || len(id) != 7 {
		return nil, errors.New("empty or invalid Id")
	}
	return &Pos{id}, nil
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
