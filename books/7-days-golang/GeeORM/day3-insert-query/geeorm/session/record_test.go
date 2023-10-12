package session

import (
	"testing"
)

var (
	user1 = &User{"A", 1}
	user2 = &User{"B", 2}
	user3 = &User{"C", 3}
)

func testInit(t *testing.T) *Session {
	t.Helper()
	s := newSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed to init test records")
	}
	return s
}

func TestSession_Insert(t *testing.T) {
	s := testInit(t)
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}

func TestSession_Find(t *testing.T) {
	s := testInit(t)
	var users []User
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query records")
	}
}
