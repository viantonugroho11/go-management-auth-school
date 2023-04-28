package major

import (
	"database/sql"
	"strings"


)

type Major struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	Table = "major"
	

	Column = []string{`def.id`, `def.name`}

	Select = `Select ` +strings.Join(Column, ", ") + ` from ` + Table + ` def `

	GroupStatement = `group by def.id `

)

func (major *Major) ScanRows(rows *sql.Rows, row *sql.Row) error {
	params := []interface{}{&major.ID, &major.Name}
	if rows != nil {
		return rows.Scan(params...)
	}
	return row.Scan(params...)
}
