package session

import (
	"geeorm/clause"
	"reflect"
)

func (s *Session) Insert(vals ...any) (int64, error) {
	recordVals := make([]any, 0)
	for _, val := range vals {
		table := s.Model(val).RefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordVals = append(recordVals, table.RecordValue(val))
	}

	s.clause.Set(clause.VALUES, recordVals...)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	res, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (s *Session) Find(vals any) error {
	dstSlice := reflect.Indirect(reflect.ValueOf(vals))
	dstType := dstSlice.Type().Elem()
	table := s.Model(reflect.New(dstType).Elem().Interface()).RefTable()

	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}

	for rows.Next() {
		dst := reflect.New(dstType).Elem()
		var fieldVals []any
		for _, name := range table.FieldNames {
			fieldVals = append(fieldVals, dst.FieldByName(name).Addr().Interface())
		}
		if err := rows.Scan(fieldVals...); err != nil {
			return err
		}
		dstSlice.Set(reflect.Append(dstSlice, dst))
	}
	return rows.Close()
}
