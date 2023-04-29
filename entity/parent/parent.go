package parent

import (
	"database/sql"
	"strings"
)

// studentEntity "go-management-auth-school/entity/student"

type Parent struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	NIK       string `json:"nik"`
	Type      string `json:"type"`
	Gender    int     `json:"gender"`
	Phone     string  `json:"phone"`
	WorkID    int     `json:"work_id"`
	WorkName  string  `json:"work_name"`
	Income    int     `json:"income"`
	StudentID string  `json:"student_id"`
	Image     string  `json:"image"`
	Student   Student `json:"student"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type Student struct {
	ID           string `json:"id"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Email        string `json:"Email"`
	Nisn         string `json:"Nisn"`
	Nis          string `json:"Nis"`
	Nik          string `json:"Nik"`
	PlaceOfBirth string `json:"PlaceOfBirth"`
	DateOfBirth  string `json:"DateOfBirth"`

	Phone    string `json:"Phone"`
	Address  string `json:"Address"`
	Gender   string `json:"gender"`
	Religion string `json:"Religion"`
	Image    string `json:"Image"`
	Status   int    `json:"Status"`

	IsActive      int `json:"IsActive"`
	ProvinceID    int `json:"ProvinceID"`
	CityID        int `json:"CityID"`
	SubdistrictID int `json:"SubdistrictID"`

	WardID int `json:"WardID"`
	Rt     int `json:"Rt"`
	Rw     int `json:"Rw"`
	// PostalCode int `json:"PostalCode"`

	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BloodType      string `json:"BloodType"`
	Disability     int    `json:"Disability"`
	DisabilityInfo string `json:"DisabilityInfo"`

	JoinDate string      `json:"JoinDate"`
	Details  interface{} `json:"Details"`
}

var (
	TableName = "parent"

	Column = []string{`def.id`, `def.first_name`, `def.last_name`, `def.nik`, `def.type`, `def.gender`, `def.phone`,
		`def.work_id`, `def.work_name`, `def.income`, `def.student_id`, `def.image`, `def.created_at`, `def.updated_at`}

	JoinColumn = []string{`LEFT JOIN student st ON st.id = def.student_id`}

	SelectParent = `SELECT ` + strings.Join(Column, `,`) + ` FROM ` + TableName + ` def`

	GroupStatement = `GROUP BY def.id`
)

func (m *Parent) ScanRows(rows *sql.Rows, row *sql.Row) error {
	parameters := []interface{}{&m.ID, &m.FirstName, &m.LastName, &m.NIK, &m.Type, &m.Gender, &m.Phone, &m.WorkID,
		&m.WorkName, &m.Income, &m.StudentID, &m.Image, &m.CreatedAt, &m.UpdatedAt}
	if rows != nil {
		return rows.Scan(parameters...)
	}
	return row.Scan(parameters...)
}
