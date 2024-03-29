package mappingstudent

import (
	"database/sql"
	classEntity "go-management-auth-school/entity/class"
	mpCourseEntity "go-management-auth-school/entity/mapping_course"
	studentEntity "go-management-auth-school/entity/student"
	teacherEntity "go-management-auth-school/entity/teacher"
	"strings"
)

type MappingStudent struct {
	ID       string                            `json:"id"`
	Identity string                         `json:"identity"`
	ClassID  string                         `json:"class_id"`
	Type     string                         `json:"type"`
	Student  studentEntity.Student          `json:"student"`
	Class    classEntity.Class              `json:"class"`
	Teacher  teacherEntity.Teacher          `json:"teacher"`
	MpCourse []mpCourseEntity.MappingCourse `json:"mp_course"`
}

type MappingStudentReq struct {
	ID       string    `json:"id"`
	IdentityID string `json:"identity"`
	ClassID  string `json:"class_id"`
	Type     string `json:"type"`
}

var (
	Table = "mapping_student"

	Column = []string{"def.id", "def.identity_id", "def.class_id", "def.type"}

	JoinGeneral = `LEFT JOIN student st ON st.id = def.identity_id
			LEFT JOIN class cl ON cl.id = def.class_id
			LEFT JOIN teacher te ON te.id = def.identity_id`

	SelectMapStudent = `SELECT ` + strings.Join(Column, `,`) + ` FROM ` + Table + ` def`

	GroupStatement = ` GROUP BY def.id`
)

func (m *MappingStudent) ScanRows(rows *sql.Rows, row *sql.Row) error {
	parameter := []interface{}{&m.ID, &m.Identity, &m.ClassID, &m.Type}
	if rows != nil {
		return rows.Scan(parameter...)
	}
	return row.Scan(parameter...)
}
