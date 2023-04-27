package parent

import (
	parentEntity "go-management-auth-school/entity/parent"
)



var (
// id VARCHAR(255) NOT NULL DEFAULT '',
//   first_name VARCHAR(255) NOT NULL DEFAULT '',
//   last_name VARCHAR(255) NOT NULL DEFAULT '',
//   type VARCHAR(255) NOT NULL DEFAULT '',

//   nik VARCHAR(255) NOT NULL DEFAULT '',
  
//   gender int NOT NULL DEFAULT 0,
//   phone VARCHAR(255) NOT NULL DEFAULT '',
//   work_id int NOT NULL DEFAULT 0,
//   work_name VARCHAR(255) NOT NULL DEFAULT '',
//   income int NOT NULL DEFAULT 0,
//   student_id VARCHAR(255) NOT NULL DEFAULT '',

//   image VARCHAR(255) NULL DEFAULT '',
	InsertParent = `INSERT INTO `+parentEntity.TableName+` (
		id,
		first_name,
		last_name,
		type,
		nik,
		gender,
		phone,
		work_id,
		work_name,
		income,
		student_id,
		image
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
		?
	)`


)