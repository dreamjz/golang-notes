package geeorm

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func openDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal("failed to connect database")
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engine := openDB(t)
	defer engine.Close()
}
