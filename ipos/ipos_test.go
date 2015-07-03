package ipos

import (
	"encoding/json"
	"fmt"
	"testing"
)

var mockPos = []byte(`{
"staffId" : "01",
"userId": "0000333",
"type": 0,
"items" : [
    {
        "itemId" : "1001",
        "qty" : 2
    },
    {
        "itemId" : "2001",
        "qty" : 1
    }
 ]
}`)

func newPosOrFatal(t *testing.T, staffId string) *Pos {
	if len(staffId) < 2 {
		var pos Pos
		err := json.Unmarshal(mockPos, &pos)
		if err != nil {
			t.Fatalf("unexpected err %v", err)
		}
		fmt.Println("newPos as", pos)
		return &pos
	}
	pos, err := NewPos(staffId)
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	return pos
}

func TestNewPos(t *testing.T) {
	pos := newPosOrFatal(t, "01")
	if pos.StaffId != "01" {
		t.Errorf("expected pos for 1 but got %v", pos.StaffId)
	}
}

func TestNewPosWithEmptyId(t *testing.T) {
	pos, err := NewPos("")
	if err == nil {
		t.Errorf("wrong or empty Id expected nil but got err:%v %#v", err, pos)
	}
}

func TestSavePos(t *testing.T) {
	pos := newPosOrFatal(t, "")

	m := NewPosManager()
	err := m.Save(pos)
	if err != nil {
		t.Errorf("unexpected err: %v", err)
	}
}

func TestDeletePos(t *testing.T) {
	pos := newPosOrFatal(t, "")

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
