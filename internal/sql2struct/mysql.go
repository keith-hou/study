package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo *DBInfo
}

type DBInfo struct {
	DBType string
	Host string
	UserName string
	Password string
	Charset string
}

type TableColumn struct {
	ColumnName string
	DataType string
	IsNullable string
	ColumnKey string
	ColumnType string
	ColumnComment string
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local"
	dns := fmt.Sprintf(s,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
		)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dns)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) GetColumns(dbname string, tableName string)([]*TableColumn, error) {
	query := "select COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT" +
		" FROM CLOUMNS where TABLE_SCHEMA=? AND TABLE_NAME= ?"
	rows, err := m.DBEngine.Query(query, dbname, tableName)
	if err != nil {
		return nil ,err
	}

	if rows == nil {
		return nil, errors.New("数据为空")
	}

	defer rows.Close()

	var columns []*TableColumn

	for rows.Next() {
		var column TableColumn
		err = rows.Scan(&column.ColumnName, &column.ColumnKey, &column.ColumnType, &column.DataType, &column.IsNullable, &column.ColumnComment)
		if err != nil {
			panic(err)
		}
		columns = append(columns, &column)
	}

	return columns, nil
}
