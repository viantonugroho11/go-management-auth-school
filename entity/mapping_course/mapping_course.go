package mapping_course

import (
	"database/sql"
	classEntity "go-management-auth-school/entity/class"
	lessonEntity "go-management-auth-school/entity/lesson"
	teacherEntity "go-management-auth-school/entity/teacher"
	"strings"
)

type MappingCourse struct {
	ID        string                `json:"id"`
	ClassID   string                `json:"class_id"`
	TeacherID string                `json:"teacher_id"`
	LessonID  string                `json:"lesson_id"`
	Class     classEntity.Class     `json:"class"`
	Teacher   teacherEntity.Teacher `json:"teacher"`
	Lesson    lessonEntity.Lesson   `json:"lesson"`
}

type MappingCourseReq struct {
	ID        string `json:"id"`
	ClassID   string `json:"class_id"`
	TeacherID string `json:"teacher_id"`
	LessonID  string `json:"lesson_id"`
}

var (
	Table = "mapping_course_teacher"

	Column = []string{"def.id", "def.class_id", "def.teacher_id", "def.lesson_id"}

	JoinGeneral = `LEFT JOIN class cl ON cl.id = def.class_id
			LEFT JOIN teacher te ON te.id = def.teacher_id
			LEFT JOIN lesson le ON le.id = def.lesson_id`

	GroupStatement = ` GROUP BY def.id`

	SelectMapCourse = `Select ` + strings.Join(Column, `,`) + ` FROM ` + Table + ` def`
)

func (m *MappingCourse) ScanRows(rows *sql.Rows, row *sql.Row) error {
	parameter := []interface{}{&m.ID, &m.ClassID, &m.TeacherID, &m.LessonID}
	if rows != nil {
		return rows.Scan(parameter...)
	}
	return row.Scan(parameter...)
}
