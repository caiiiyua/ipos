package ipos

import "testing"

func newPosOrFatal(t *testing.T, id string) *Pos {
	pos, err := NewPos(id)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	return pos
}

func TestNewPos(t *testing.T) {
	pos := newPosOrFatal(t, "0000001")
	if pos.Id != "0000001" {
		t.Errorf("expected pos for 1 but got %v", pos.Id)
	}
}

func TestNewPosWithEmptyId(t *testing.T) {
	pos, err := NewPos("")
	if err == nil {
		t.Errorf("wrong or empty Id expected nil but got err:%v %#v", err, pos)
	}
}

func TestSavePos(t *testing.T) {
	pos := newPosOrFatal(t, "0000001")

	m := NewPosManager()
	err := m.Save(pos)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
}

func TestDeletePos(t *testing.T) {
	pos := newPosOrFatal(t, "0000001")

	m := NewPosManager()
	err := m.Delete(pos)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
}

func TestListPos(t *testing.T) {
	m := NewPosManager()
	poses, err := m.List()
	if err != nil {
		t.Errorf("unexpected err %v with %v pos", err, len(poses))
	}
}
