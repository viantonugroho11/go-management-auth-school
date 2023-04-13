package user

import (
	"database/sql"
	studentEntity "go-management-auth-school/entity/student"
	"strings"
)

// id VARCHAR(255) NOT NULL DEFAULT '',
//   username VARCHAR(255) NOT NULL DEFAULT '',
//   password VARCHAR(255) NOT NULL DEFAULT '',
//   identify_id VARCHAR(255) NOT NULL DEFAULT '',
//   permission INT NOT NULL DEFAULT 0,
//   last_login DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00',
//   device_id VARCHAR(255) NOT NULL DEFAULT '',
//   status boolean NOT NULL DEFAULT 1,
//   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//   deleted_at TIMESTAMP NULL DEFAULT NULL,
//   PRIMARY KEY (id)
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IdentityID string `json:"identify_id"`
	Permission int `json:"permission"`
	LastLogin string `json:"last_login"`
	DeviceID string `json:"device_id"`
	Status bool `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt sql.NullString `json:"deleted_at"`
}

var (
	Table = `user`
	Column = []string{`def.id`,`def.username`,`def.password`,`def.identity_id`,`def.permission`,`def.last_login`,
	`def.device_id`,`def.status`,`def.created_at`,`def.updated_at`,`def.deleted_at`}

	ColumnStudent = []string{`st.id`,`st.first_name`,`st.last_name`,`st.email`,`st.nisn`,`st.nis`,
	`st.nik`,`st.place_of_birth`,`st.date_of_birth`,`st.phone`,`st.address`,`st.gender`,
	`st.religion`,`st.image`,`st.status`,`st.is_active`,
	`st.deleted_at`}

	SelectUser = `SELECT ` + strings.Join(Column, `,`) + ` FROM ` + Table + ` def`

	JoinStatment = `LEFT JOIN ` + studentEntity.Table + ` st ON st.id = def.identity_id AND def.permission = 1 AND def.status = 1 AND def.deleted_at IS NULL`
)

func (u *User) ScanRows(rows *sql.Rows, row *sql.Row) error {
	parameter := []interface{}{&u.ID, &u.Username, &u.Password, &u.IdentityID, &u.Permission, &u.LastLogin,
		&u.DeviceID, &u.Status, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt}
	if rows != nil {
		return rows.Scan(parameter...)
	}
	return row.Scan(parameter...)
}

