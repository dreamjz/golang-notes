package geeorm

import (
	"database/sql"
	"geeorm/log"
	"geeorm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (*Engine, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Send ping to make sure the database connection is alive
	if err := db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}

	e := &Engine{db: db}
	log.Info("Connect database success")

	return e, nil
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database connection")
		return
	}
	log.Info("Close database connection success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
