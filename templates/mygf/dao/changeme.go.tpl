package dao

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// {{.StructName}}Dao is the data access object for table report.
type {{.StructName}}Dao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns {{.StructName}}Columns // columns contains all the column names of Table for convenient usage.
}

// {{.StructName}}Columns defines and stores column names for table report.
type {{.StructName}}Columns struct {
	{{- range $idx, $field := .FieldMap}}
	{{$field.Name}} string // {{$field.Desc}}
	{{- end}}
}

//  reportColumns holds the columns for table report.
var reportColumns = {{.StructName}}Columns{
	{{- range $idx, $field := .FieldMap}}
	{{$field.Name}}: "{{$field.Orm}}" ,
	{{- end}}
}

var (
	// {{.StructName}} is globally public accessible object for table report operations.
	{{.StructName}} = New{{.StructName}}Dao()
)

// New{{.StructName}}Dao creates and returns a new DAO object for table data access.
func New{{.StructName}}Dao() *{{.StructName}}Dao {
	return &{{.StructName}}Dao{
		group:   "default",
		table:   "report",
		columns: reportColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *{{.StructName}}Dao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *{{.StructName}}Dao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *{{.StructName}}Dao) Columns() {{.StructName}}Columns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *{{.StructName}}Dao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *{{.StructName}}Dao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *{{.StructName}}Dao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
