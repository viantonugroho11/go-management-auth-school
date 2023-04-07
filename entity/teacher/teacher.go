package teacher

type Teacher struct {
	ID string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Nik string `json:"nik"`
	PlaceOfBirth string `json:"place_of_birth"`
	DateOfBirth string `json:"date_of_birth"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	Gender int `json:"gender"`
	Religion string `json:"religion"`
	Image string `json:"image"`
	Status int `json:"status"`
	IsActive int `json:"is_active"`
	ProvinceID int `json:"province_id"`
	CityID int `json:"city_id"`
	SubdistrictID int `json:"subdistrict_id"`
	WardID int `json:"ward_id"`
	RT int `json:"rt"`
	RW int `json:"rw"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}