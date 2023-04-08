package student

import (
	studentEntity "go-management-auth-school/entity/student"
)

var (
	InsertStudent = `INSERT INTO `+studentEntity.Table+` (
		id, 
		first_name, 
		last_name, 
		email, 
		nisn, 
		nis, 
		nik,
		place_of_birth,
		date_of_birth,
		phone,
		address,
		gender,
		religion,
		image,
		status,
		is_active,
		province_id,
		city_id,
		subdistrict_id,
		ward_id,
		rt,
		rw,
		height,
		weight,
		blood_type,
		disability,
		disability_info,
		join_date,
		details
	) VALUES (
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?
	)`

	
)