package lesson

import (
	"database/sql"
	"strings"
)

type Lesson struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var (
	Table = "lesson"

	Columns = []string{`def.id`, `def.name`, `def.type`}

	SelectStatment = `Select ` + strings.Join(Columns, ",") + ` From ` + Table + ` def `

	GroupStatement = ` Group By def.id `
)

func (lesson *Lesson) ScanRows(rows *sql.Rows, row *sql.Row) error {
	parameter := []interface{}{&lesson.ID, &lesson.Name, &lesson.Type}
	if rows != nil {
		return rows.Scan(parameter...)
	}
	return row.Scan(parameter...)
}
