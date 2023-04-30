package verify_token

import (
	"database/sql"
	"strings"
)


type VerifyToken struct {
	ID string `json:"id"`
	Identity string `json:"identity"`
	Token string `json:"token"`
	ExpiredAt	string `json:"expired_at"`
}

var (
	TableName = "verify_token"

	ColumnVerifyToken = []string{`def.id`, `def.identity_id`, `def.token`, `def.expired_at`}

	SelectStatement = `SELECT ` + strings.Join(ColumnVerifyToken, ", ") + ` FROM ` + TableName + ` def`

	GroupStatement = ` GROUP BY def.id `
)

func (m *VerifyToken) ScanRows(rows *sql.Rows, row *sql.Row) error {
	parameter := []interface{}{&m.ID, &m.Identity, &m.Token, &m.ExpiredAt}
	if rows != nil {
		return rows.Scan(parameter...)
	}
	return row.Scan(parameter...)
}