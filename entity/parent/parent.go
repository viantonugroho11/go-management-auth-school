package parent


type Parent struct {
	ID string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Type string `json:"type"`

	Gender int `json:"gender"`
	Phone string `json:"phone"`
	WorkID int `json:"work_id"`
	WorkName string `json:"work_name"`
	Income int `json:"income"`
	StudentID string `json:"student_id"`
	Image string `json:"image"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}