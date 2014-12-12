// Code generated by ModelQ, 2014-12-12 22:55
// COLUMNS.go contains model for the database table [information_schema.COLUMNS]

package meta

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/mijia/modelq/gmq"
	"strings"
)

type Columns struct {
	TableCatalog           string `json:"TABLE_CATALOG"`
	TableSchema            string `json:"TABLE_SCHEMA"`
	TableName              string `json:"TABLE_NAME"`
	ColumnName             string `json:"COLUMN_NAME"`
	OrdinalPosition        int64  `json:"ORDINAL_POSITION"`
	ColumnDefault          string `json:"COLUMN_DEFAULT"`
	IsNullable             string `json:"IS_NULLABLE"`
	DataType               string `json:"DATA_TYPE"`
	CharacterMaximumLength int64  `json:"CHARACTER_MAXIMUM_LENGTH"`
	CharacterOctetLength   int64  `json:"CHARACTER_OCTET_LENGTH"`
	NumericPrecision       int64  `json:"NUMERIC_PRECISION"`
	NumericScale           int64  `json:"NUMERIC_SCALE"`
	DatetimePrecision      int64  `json:"DATETIME_PRECISION"`
	CharacterSetName       string `json:"CHARACTER_SET_NAME"`
	CollationName          string `json:"COLLATION_NAME"`
	ColumnType             string `json:"COLUMN_TYPE"`
	ColumnKey              string `json:"COLUMN_KEY"`
	Extra                  string `json:"EXTRA"`
	Privileges             string `json:"PRIVILEGES"`
	ColumnComment          string `json:"COLUMN_COMMENT"`
}

// Start of the Columns APIs.

func (obj Columns) String() string {
	if data, err := json.Marshal(obj); err != nil {
		return fmt.Sprintf("<Columns>")
	} else {
		return string(data)
	}
}

func (obj Columns) Insert(dbtx gmq.DbTx) (Columns, error) {
	_, err := ColumnsObjs.Insert(obj).Run(dbtx)
	return obj, err
}

func (obj Columns) Update(dbtx gmq.DbTx) (int64, error) {
	return 0, ErrNoPrimaryKeyDefined
}

func (obj Columns) Delete(dbtx gmq.DbTx) (int64, error) {
	return 0, ErrNoPrimaryKeyDefined
}

// Start of the inner Query Api

type _ColumnsQuery struct {
	gmq.Query
}

func (q _ColumnsQuery) Where(f gmq.Filter) _ColumnsQuery {
	q.Query = q.Query.Where(f)
	return q
}

func (q _ColumnsQuery) OrderBy(by ...string) _ColumnsQuery {
	tBy := make([]string, 0, len(by))
	for _, b := range by {
		sortDir := ""
		if b[0] == '-' || b[0] == '+' {
			sortDir = string(b[0])
			b = b[1:]
		}
		if col, ok := ColumnsObjs.fcMap[b]; ok {
			tBy = append(tBy, sortDir+col)
		}
	}
	q.Query = q.Query.OrderBy(tBy...)
	return q
}

func (q _ColumnsQuery) GroupBy(by ...string) _ColumnsQuery {
	tBy := make([]string, 0, len(by))
	for _, b := range by {
		if col, ok := ColumnsObjs.fcMap[b]; ok {
			tBy = append(tBy, col)
		}
	}
	q.Query = q.Query.GroupBy(tBy...)
	return q
}

func (q _ColumnsQuery) Limit(offsets ...int64) _ColumnsQuery {
	q.Query = q.Query.Limit(offsets...)
	return q
}

func (q _ColumnsQuery) Page(number, size int) _ColumnsQuery {
	q.Query = q.Query.Page(number, size)
	return q
}

func (q _ColumnsQuery) Run(dbtx gmq.DbTx) (sql.Result, error) {
	return q.Query.Exec(dbtx)
}

type ColumnsRowVisitor func(obj Columns) bool

func (q _ColumnsQuery) Iterate(dbtx gmq.DbTx, functor ColumnsRowVisitor) error {
	return q.Query.SelectList(dbtx, func(columns []gmq.Column, rb []sql.RawBytes) bool {
		obj := ColumnsObjs.toColumns(columns, rb)
		return functor(obj)
	})
}

func (q _ColumnsQuery) One(dbtx gmq.DbTx) (Columns, error) {
	var obj Columns
	err := q.Query.SelectOne(dbtx, func(columns []gmq.Column, rb []sql.RawBytes) bool {
		obj = ColumnsObjs.toColumns(columns, rb)
		return true
	})
	return obj, err
}

func (q _ColumnsQuery) List(dbtx gmq.DbTx) ([]Columns, error) {
	result := make([]Columns, 0, 10)
	err := q.Query.SelectList(dbtx, func(columns []gmq.Column, rb []sql.RawBytes) bool {
		obj := ColumnsObjs.toColumns(columns, rb)
		result = append(result, obj)
		return true
	})
	return result, err
}

// Start of the model facade Apis.

type _ColumnsObjs struct {
	fcMap map[string]string
}

func (o _ColumnsObjs) Names() (string, string) { return "COLUMNS", "Columns" }

func (o _ColumnsObjs) Select(fields ...string) _ColumnsQuery {
	q := _ColumnsQuery{}
	if len(fields) == 0 {
		fields = []string{"TableCatalog", "TableSchema", "TableName", "ColumnName", "OrdinalPosition", "ColumnDefault", "IsNullable", "DataType", "CharacterMaximumLength", "CharacterOctetLength", "NumericPrecision", "NumericScale", "DatetimePrecision", "CharacterSetName", "CollationName", "ColumnType", "ColumnKey", "Extra", "Privileges", "ColumnComment"}
	}
	q.Query = gmq.Select(o, o.columns(fields...))
	return q
}

func (o _ColumnsObjs) Insert(obj Columns) _ColumnsQuery {
	q := _ColumnsQuery{}
	q.Query = gmq.Insert(o, o.columnsWithData(obj, "TableCatalog", "TableSchema", "TableName", "ColumnName", "OrdinalPosition", "ColumnDefault", "IsNullable", "DataType", "CharacterMaximumLength", "CharacterOctetLength", "NumericPrecision", "NumericScale", "DatetimePrecision", "CharacterSetName", "CollationName", "ColumnType", "ColumnKey", "Extra", "Privileges", "ColumnComment"))
	return q
}

func (o _ColumnsObjs) Update(obj Columns, fields ...string) _ColumnsQuery {
	q := _ColumnsQuery{}
	q.Query = gmq.Update(o, o.columnsWithData(obj, fields...))
	return q
}

func (o _ColumnsObjs) Delete() _ColumnsQuery {
	q := _ColumnsQuery{}
	q.Query = gmq.Delete(o)
	return q
}

///// Managed Objects Filters definition

func (o _ColumnsObjs) FilterTableCatalog(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("TABLE_CATALOG", op, params...)
}

func (o _ColumnsObjs) FilterTableSchema(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("TABLE_SCHEMA", op, params...)
}

func (o _ColumnsObjs) FilterTableName(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("TABLE_NAME", op, params...)
}

func (o _ColumnsObjs) FilterColumnName(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("COLUMN_NAME", op, params...)
}

func (o _ColumnsObjs) FilterOrdinalPosition(op string, p int64, ps ...int64) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("ORDINAL_POSITION", op, params...)
}

func (o _ColumnsObjs) FilterColumnDefault(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("COLUMN_DEFAULT", op, params...)
}

func (o _ColumnsObjs) FilterIsNullable(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("IS_NULLABLE", op, params...)
}

func (o _ColumnsObjs) FilterDataType(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("DATA_TYPE", op, params...)
}

func (o _ColumnsObjs) FilterCharacterMaximumLength(op string, p int64, ps ...int64) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("CHARACTER_MAXIMUM_LENGTH", op, params...)
}

func (o _ColumnsObjs) FilterCharacterOctetLength(op string, p int64, ps ...int64) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("CHARACTER_OCTET_LENGTH", op, params...)
}

func (o _ColumnsObjs) FilterNumericPrecision(op string, p int64, ps ...int64) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("NUMERIC_PRECISION", op, params...)
}

func (o _ColumnsObjs) FilterNumericScale(op string, p int64, ps ...int64) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("NUMERIC_SCALE", op, params...)
}

func (o _ColumnsObjs) FilterDatetimePrecision(op string, p int64, ps ...int64) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("DATETIME_PRECISION", op, params...)
}

func (o _ColumnsObjs) FilterCharacterSetName(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("CHARACTER_SET_NAME", op, params...)
}

func (o _ColumnsObjs) FilterCollationName(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("COLLATION_NAME", op, params...)
}

func (o _ColumnsObjs) FilterColumnType(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("COLUMN_TYPE", op, params...)
}

func (o _ColumnsObjs) FilterColumnKey(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("COLUMN_KEY", op, params...)
}

func (o _ColumnsObjs) FilterExtra(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("EXTRA", op, params...)
}

func (o _ColumnsObjs) FilterPrivileges(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("PRIVILEGES", op, params...)
}

func (o _ColumnsObjs) FilterColumnComment(op string, p string, ps ...string) gmq.Filter {
	params := make([]interface{}, 1+len(ps))
	params[0] = p
	for i := range ps {
		params[i+1] = ps[i]
	}
	return o.newFilter("COLUMN_COMMENT", op, params...)
}

///// Managed Objects Columns definition

func (o _ColumnsObjs) ColumnTableCatalog(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"TABLE_CATALOG", value}
}

func (o _ColumnsObjs) ColumnTableSchema(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"TABLE_SCHEMA", value}
}

func (o _ColumnsObjs) ColumnTableName(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"TABLE_NAME", value}
}

func (o _ColumnsObjs) ColumnColumnName(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"COLUMN_NAME", value}
}

func (o _ColumnsObjs) ColumnOrdinalPosition(p ...int64) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"ORDINAL_POSITION", value}
}

func (o _ColumnsObjs) ColumnColumnDefault(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"COLUMN_DEFAULT", value}
}

func (o _ColumnsObjs) ColumnIsNullable(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"IS_NULLABLE", value}
}

func (o _ColumnsObjs) ColumnDataType(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"DATA_TYPE", value}
}

func (o _ColumnsObjs) ColumnCharacterMaximumLength(p ...int64) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"CHARACTER_MAXIMUM_LENGTH", value}
}

func (o _ColumnsObjs) ColumnCharacterOctetLength(p ...int64) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"CHARACTER_OCTET_LENGTH", value}
}

func (o _ColumnsObjs) ColumnNumericPrecision(p ...int64) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"NUMERIC_PRECISION", value}
}

func (o _ColumnsObjs) ColumnNumericScale(p ...int64) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"NUMERIC_SCALE", value}
}

func (o _ColumnsObjs) ColumnDatetimePrecision(p ...int64) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"DATETIME_PRECISION", value}
}

func (o _ColumnsObjs) ColumnCharacterSetName(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"CHARACTER_SET_NAME", value}
}

func (o _ColumnsObjs) ColumnCollationName(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"COLLATION_NAME", value}
}

func (o _ColumnsObjs) ColumnColumnType(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"COLUMN_TYPE", value}
}

func (o _ColumnsObjs) ColumnColumnKey(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"COLUMN_KEY", value}
}

func (o _ColumnsObjs) ColumnExtra(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"EXTRA", value}
}

func (o _ColumnsObjs) ColumnPrivileges(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"PRIVILEGES", value}
}

func (o _ColumnsObjs) ColumnColumnComment(p ...string) gmq.Column {
	var value interface{}
	if len(p) > 0 {
		value = p[0]
	}
	return gmq.Column{"COLUMN_COMMENT", value}
}

////// Internal helper funcs

func (o _ColumnsObjs) newFilter(name, op string, params ...interface{}) gmq.Filter {
	if strings.ToUpper(op) == "IN" {
		return gmq.InFilter(name, params)
	}
	return gmq.UnitFilter(name, op, params[0])
}

func (o _ColumnsObjs) toColumns(columns []gmq.Column, rb []sql.RawBytes) Columns {
	obj := Columns{}
	if len(columns) == len(rb) {
		for i := range columns {
			switch columns[i].Name {
			case "TABLE_CATALOG":
				obj.TableCatalog = gmq.AsString(rb[i])
			case "TABLE_SCHEMA":
				obj.TableSchema = gmq.AsString(rb[i])
			case "TABLE_NAME":
				obj.TableName = gmq.AsString(rb[i])
			case "COLUMN_NAME":
				obj.ColumnName = gmq.AsString(rb[i])
			case "ORDINAL_POSITION":
				obj.OrdinalPosition = gmq.AsInt64(rb[i])
			case "COLUMN_DEFAULT":
				obj.ColumnDefault = gmq.AsString(rb[i])
			case "IS_NULLABLE":
				obj.IsNullable = gmq.AsString(rb[i])
			case "DATA_TYPE":
				obj.DataType = gmq.AsString(rb[i])
			case "CHARACTER_MAXIMUM_LENGTH":
				obj.CharacterMaximumLength = gmq.AsInt64(rb[i])
			case "CHARACTER_OCTET_LENGTH":
				obj.CharacterOctetLength = gmq.AsInt64(rb[i])
			case "NUMERIC_PRECISION":
				obj.NumericPrecision = gmq.AsInt64(rb[i])
			case "NUMERIC_SCALE":
				obj.NumericScale = gmq.AsInt64(rb[i])
			case "DATETIME_PRECISION":
				obj.DatetimePrecision = gmq.AsInt64(rb[i])
			case "CHARACTER_SET_NAME":
				obj.CharacterSetName = gmq.AsString(rb[i])
			case "COLLATION_NAME":
				obj.CollationName = gmq.AsString(rb[i])
			case "COLUMN_TYPE":
				obj.ColumnType = gmq.AsString(rb[i])
			case "COLUMN_KEY":
				obj.ColumnKey = gmq.AsString(rb[i])
			case "EXTRA":
				obj.Extra = gmq.AsString(rb[i])
			case "PRIVILEGES":
				obj.Privileges = gmq.AsString(rb[i])
			case "COLUMN_COMMENT":
				obj.ColumnComment = gmq.AsString(rb[i])
			}
		}
	}
	return obj
}

func (o _ColumnsObjs) columns(fields ...string) []gmq.Column {
	data := make([]gmq.Column, 0, len(fields))
	for _, f := range fields {
		switch f {
		case "TableCatalog":
			data = append(data, o.ColumnTableCatalog())
		case "TableSchema":
			data = append(data, o.ColumnTableSchema())
		case "TableName":
			data = append(data, o.ColumnTableName())
		case "ColumnName":
			data = append(data, o.ColumnColumnName())
		case "OrdinalPosition":
			data = append(data, o.ColumnOrdinalPosition())
		case "ColumnDefault":
			data = append(data, o.ColumnColumnDefault())
		case "IsNullable":
			data = append(data, o.ColumnIsNullable())
		case "DataType":
			data = append(data, o.ColumnDataType())
		case "CharacterMaximumLength":
			data = append(data, o.ColumnCharacterMaximumLength())
		case "CharacterOctetLength":
			data = append(data, o.ColumnCharacterOctetLength())
		case "NumericPrecision":
			data = append(data, o.ColumnNumericPrecision())
		case "NumericScale":
			data = append(data, o.ColumnNumericScale())
		case "DatetimePrecision":
			data = append(data, o.ColumnDatetimePrecision())
		case "CharacterSetName":
			data = append(data, o.ColumnCharacterSetName())
		case "CollationName":
			data = append(data, o.ColumnCollationName())
		case "ColumnType":
			data = append(data, o.ColumnColumnType())
		case "ColumnKey":
			data = append(data, o.ColumnColumnKey())
		case "Extra":
			data = append(data, o.ColumnExtra())
		case "Privileges":
			data = append(data, o.ColumnPrivileges())
		case "ColumnComment":
			data = append(data, o.ColumnColumnComment())
		}
	}
	return data
}

func (o _ColumnsObjs) columnsWithData(obj Columns, fields ...string) []gmq.Column {
	data := make([]gmq.Column, 0, len(fields))
	for _, f := range fields {
		switch f {
		case "TableCatalog":
			data = append(data, o.ColumnTableCatalog(obj.TableCatalog))
		case "TableSchema":
			data = append(data, o.ColumnTableSchema(obj.TableSchema))
		case "TableName":
			data = append(data, o.ColumnTableName(obj.TableName))
		case "ColumnName":
			data = append(data, o.ColumnColumnName(obj.ColumnName))
		case "OrdinalPosition":
			data = append(data, o.ColumnOrdinalPosition(obj.OrdinalPosition))
		case "ColumnDefault":
			data = append(data, o.ColumnColumnDefault(obj.ColumnDefault))
		case "IsNullable":
			data = append(data, o.ColumnIsNullable(obj.IsNullable))
		case "DataType":
			data = append(data, o.ColumnDataType(obj.DataType))
		case "CharacterMaximumLength":
			data = append(data, o.ColumnCharacterMaximumLength(obj.CharacterMaximumLength))
		case "CharacterOctetLength":
			data = append(data, o.ColumnCharacterOctetLength(obj.CharacterOctetLength))
		case "NumericPrecision":
			data = append(data, o.ColumnNumericPrecision(obj.NumericPrecision))
		case "NumericScale":
			data = append(data, o.ColumnNumericScale(obj.NumericScale))
		case "DatetimePrecision":
			data = append(data, o.ColumnDatetimePrecision(obj.DatetimePrecision))
		case "CharacterSetName":
			data = append(data, o.ColumnCharacterSetName(obj.CharacterSetName))
		case "CollationName":
			data = append(data, o.ColumnCollationName(obj.CollationName))
		case "ColumnType":
			data = append(data, o.ColumnColumnType(obj.ColumnType))
		case "ColumnKey":
			data = append(data, o.ColumnColumnKey(obj.ColumnKey))
		case "Extra":
			data = append(data, o.ColumnExtra(obj.Extra))
		case "Privileges":
			data = append(data, o.ColumnPrivileges(obj.Privileges))
		case "ColumnComment":
			data = append(data, o.ColumnColumnComment(obj.ColumnComment))
		}
	}
	return data
}

var ColumnsObjs _ColumnsObjs

func init() {
	ColumnsObjs.fcMap = map[string]string{
		"TableCatalog":           "TABLE_CATALOG",
		"TableSchema":            "TABLE_SCHEMA",
		"TableName":              "TABLE_NAME",
		"ColumnName":             "COLUMN_NAME",
		"OrdinalPosition":        "ORDINAL_POSITION",
		"ColumnDefault":          "COLUMN_DEFAULT",
		"IsNullable":             "IS_NULLABLE",
		"DataType":               "DATA_TYPE",
		"CharacterMaximumLength": "CHARACTER_MAXIMUM_LENGTH",
		"CharacterOctetLength":   "CHARACTER_OCTET_LENGTH",
		"NumericPrecision":       "NUMERIC_PRECISION",
		"NumericScale":           "NUMERIC_SCALE",
		"DatetimePrecision":      "DATETIME_PRECISION",
		"CharacterSetName":       "CHARACTER_SET_NAME",
		"CollationName":          "COLLATION_NAME",
		"ColumnType":             "COLUMN_TYPE",
		"ColumnKey":              "COLUMN_KEY",
		"Extra":                  "EXTRA",
		"Privileges":             "PRIVILEGES",
		"ColumnComment":          "COLUMN_COMMENT",
	}
}
