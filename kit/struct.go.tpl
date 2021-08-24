package {{.Models}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
    {{range .Imports}}"{{.}}"{{end}}
    "github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql"
)
{{end}}

{{range .Tables}}
var {{Mapper .Name}}Model = &{{Mapper .Name}}{}

func init() {
	// TODO need packaging
	{{Mapper .Name}}Model.Dao, _ = xorm.NewEngine("mysql", "root:12345678@/company?charset=utf8mb4")
}

type {{Mapper .Name}} struct {
    tableName string `xorm:"-"`
    Dao       *xorm.Engine
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}}	{{Type $col}} {{Tag $table $col}}
{{end}}
}

func (m *{{Mapper .Name}}) TableName() string {
    if m.tableName == "" {
	    return "{{.Name}}"
    }
    return m.tableName
}

func (m *{{Mapper .Name}}) Insert(r {{Mapper .Name}}) int64 {
	affected, _ := m.Dao.Insert(r)
	return affected
}

func (m *{{Mapper .Name}}) Query(r interface{}) map[string][]byte {
	res, _ := m.Dao.Query(r)
	for _, v := range res {
		return v
	}
	return nil
}

func (m *{{Mapper .Name}}) Exec(r interface{}) int64 {
	res, _ := m.Dao.Exec(r)
	affected, _ := res.RowsAffected()
	return affected
}

{{end}}