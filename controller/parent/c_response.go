package parent

import (
	// parentEntity "go-management-auth-school/entity/parent"
)

type ParentResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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
