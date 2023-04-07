package student

import (
	"database/sql"
	"strings"
)


type Student struct {
	
	ID string `json:"id"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Email string `json:"Email"`
	Nisn string `json:"Nisn"`
	Nis string `json:"Nis"`
	Nik string `json:"Nik"`
	PlaceOfBirth string `json:"PlaceOfBirth"`
	DateOfBirth string `json:"DateOfBirth"`

	Phone string `json:"Phone"`
	Address string `json:"Address"`
	Gender string `json:"gender"`
	Religion string `json:"Religion"`
	Image string `json:"Image"`
	Status int `json:"Status"`

	IsActive int `json:"IsActive"`
	ProvinceID int `json:"ProvinceID"`
	CityID int `json:"CityID"`
	SubdistrictID int `json:"SubdistrictID"`

	WardID int `json:"WardID"`
	Rt int `json:"Rt"`
	Rw int `json:"Rw"`
	// PostalCode int `json:"PostalCode"`

	Height int `json:"height"`
	Weight int `json:"weight"`
	BloodType string `json:"BloodType"`
	Disability int `json:"Disability"`
	DisabilityInfo string `json:"DisabilityInfo"`

	JoinDate string `json:"JoinDate"`
	Details interface{} `json:"Details"`
	CreatedAt string `json:"CreatedAt"`

	UpdatedAt string `json:"UpdatedAt"`
	DeletedAt string `json:"DeletedAt"`
}

var (
	Table = `student`
	Column = []string{`def.id`,`def.first_name`,`def.last_name`,`def.email`,`def.nisn`,`def.nis`,
	`def.nik`,`def.place_of_birth`,`def.date_of_birth`,`def.phone`,`def.address`,`def.gender`,
	`def.religion`,`def.image`,`def.status`,`def.is_active`,`def.province_id`,`def.city_id`,
	`def.subdistrict_id`,`def.ward_id`,`def.rt`,`def.rw`,`def.height`,`def.weight`,`def.blood_type`,
	`def.disability`,`def.disability_info`,`def.join_date`,`def.details`,`def.created_at`,`def.updated_at`,
	`def.deleted_at`}

	SelectUser = `SELECT ` + strings.Join(Column, `,`) + ` FROM ` + Table + ` def`
)

func (m *Student) ScanRows(rows *sql.Rows, row *sql.Row) error{
	parameter := []interface{}{&m.ID, &m.FirstName, &m.LastName, &m.Email, &m.Nisn, &m.Nis,
		&m.Nik, &m.PlaceOfBirth, &m.DateOfBirth, &m.Phone, &m.Address, &m.Gender,
		&m.Religion, &m.Image, &m.Status, &m.IsActive, &m.ProvinceID, &m.CityID,
		&m.SubdistrictID, &m.WardID, &m.Rt, &m.Rw, &m.Height, &m.Weight, &m.BloodType,
		&m.Disability, &m.DisabilityInfo, &m.JoinDate, &m.Details, &m.CreatedAt,
		&m.UpdatedAt, &m.DeletedAt}
	if rows != nil {
		return rows.Scan(parameter...)
	}
	return row.Scan(parameter...)
}