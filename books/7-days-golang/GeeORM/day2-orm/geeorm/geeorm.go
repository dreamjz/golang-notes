package geeorm

import (
	"database/sql"
	"errors"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/session"
)

var (
	DialectNotFoundErr = errors.New("dialect not found")
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (*Engine, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Send ping to make sure the database connection is alive
	if err = db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}

	// make sure the specific dialect exists
	d, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return nil, DialectNotFoundErr
	}

	e := &Engine{db: db, dialect: d}
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
	return session.New(engine.db, engine.dialect)
}
