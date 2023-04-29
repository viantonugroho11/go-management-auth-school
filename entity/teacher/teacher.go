package teacher

import (
	"database/sql"
	"strings"
)

type Teacher struct {
	ID            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Nik           string `json:"nik"`
	PlaceOfBirth  string `json:"place_of_birth"`
	DateOfBirth   string `json:"date_of_birth"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	Gender        int    `json:"gender"`
	Religion      string `json:"religion"`
	Image         string `json:"image"`
	Status        int    `json:"status"`
	IsActive      int    `json:"is_active"`
	ProvinceID    int    `json:"province_id"`
	CityID        int    `json:"city_id"`
	SubdistrictID int    `json:"subdistrict_id"`
	WardID        int    `json:"ward_id"`
	RT            int    `json:"rt"`
	RW            int    `json:"rw"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

var (
	Table = "teacher"

	Column = []string{`def.id`, `def.first_name`, `def.last_name`, `def.email`, `def.nik`, 
	`def.place_of_birth`, `def.date_of_birth`, `def.phone`, `def.address`, `def.gender`, 
	`def.religion`, `def.image`, `def.status`, `def.is_active`, `def.province_id`, `def.city_id`, 
	`def.subdistrict_id`, `def.ward_id`, `def.rt`, `def.rw`, `def.created_at`, `def.updated_at`}

	SelectStatement = `SELECT ` + strings.Join(Column, ",") + ` FROM ` + Table + ` def `

	GroupStatement = ` GROUP BY def.id `
)

func (m *Teacher) ScanRows(rows *sql.Rows, row *sql.Row) error {
	parameter := []interface{}{&m.ID, &m.FirstName, &m.LastName, &m.Email, &m.Nik, &m.PlaceOfBirth,
	&m.DateOfBirth, &m.Phone, &m.Address, &m.ProvinceID, &m.CityID, &m.SubdistrictID, &m.WardID, &m.RT, &m.RW,
&m.CreatedAt, &m.UpdatedAt}
	if rows != nil {
		return rows.Scan(parameter...)
	}
	return row.Scan(parameter...)
}
