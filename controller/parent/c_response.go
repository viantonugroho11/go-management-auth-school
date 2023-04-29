package parent

import (
	parentEntity "go-management-auth-school/entity/parent"
	studentEntity "go-management-auth-school/entity/student"
	helperStr "go-management-auth-school/helper/str"
)

type ParentResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Type      string `json:"type"`
	Gender    int    `json:"gender"`
	Phone     string `json:"phone"`
	WorkID    int    `json:"work_id"`
	WorkName  string `json:"work_name"`
	Income    int    `json:"income"`
	StudentID string `json:"student_id"`
	Student   string `json:"student"`
	Image     string `json:"image"`
}

func FromServices(res []parentEntity.Parent) (data []ParentResponse) {
	for _, v := range res {
		data = append(data, FromService(v))
	}
	return
}

func FromService(res parentEntity.Parent) (data ParentResponse) {
	data = ParentResponse{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		FullName:  helperStr.GetFullNameParent(res),
		Type:      res.Type,
		Gender:    res.Gender,
		Phone:     res.Phone,
		WorkID:    res.WorkID,
		WorkName:  res.WorkName,
		Income:    res.Income,
		StudentID: res.StudentID,
		Student: helperStr.GetFullNameStudent(studentEntity.Student{
			FirstName: res.Student.FirstName,
			LastName:  res.Student.LastName,
		}),
		Image: res.Image,
	}
	return
}
