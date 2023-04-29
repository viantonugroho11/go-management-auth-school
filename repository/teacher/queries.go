package teacher

import (
	teacherEntity "go-management-auth-school/entity/teacher"
)

var (
	InsertTeacher = `INSERT INTO ` +teacherEntity.Table+ `
		(id,first_name,last_name,email,nik,place_of_birth,date_of_birth,phone,address,gender,
			religion,image,status,is_active,province_id,city_id,
			subdistrict_id,ward_id,rt,rw) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
)