package class

import (
	"database/sql"
	majorEntity "go-management-auth-school/entity/major"
	"strings"
)

type Class struct {
	ID      int               `json:"id"`
	Name    string            `json:"name"`
	MajorID int               `json:"major_id"`
	Major   majorEntity.Major `json:"major"`
	Level   string            `json:"level"`
}


var (
	Table = "class"

	Columns = []string{`def.id`, `def.name`, `def.major_id`, `def.level`}

	SelectStatment = `Select ` + "`" + strings.Join(Columns, "` , `") + "`" + ` From ` + Table + ` def `

	GroupStatement = ` Group By def.id `
)

func (class *Class) ScanRows(rows *sql.Rows, row *sql.Row) error {
	parameter := []interface{}{&class.ID, &class.Name, &class.MajorID, &class.Level}
	if rows != nil {
		return rows.Scan(parameter...)
	}
	return row.Scan(parameter...)
}