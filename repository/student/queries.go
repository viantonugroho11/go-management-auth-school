package student

import (
	studentEntity "go-management-auth-school/entity/student"
)

var (
	// Query Insert
	// id VARCHAR(255) NOT NULL DEFAULT '',
  // first_name VARCHAR(255) NOT NULL DEFAULT '',
  // last_name VARCHAR(255) NOT NULL DEFAULT '',
  // email VARCHAR(255) NOT NULL DEFAULT '',

  // nisn VARCHAR(255) NOT NULL DEFAULT '',
  // nis VARCHAR(255) NOT NULL DEFAULT '',
  // nik VARCHAR(255) NOT NULL DEFAULT '',

  // place_of_birth VARCHAR(255) NOT NULL DEFAULT '',
  // date_of_birth DATE NOT NULL DEFAULT '0000-00-00',

  // phone VARCHAR(255) NOT NULL DEFAULT '',
  // address VARCHAR(255) NOT NULL DEFAULT '',
  // gender ENUM('male','female') NOT NULL DEFAULT 'male',
  // religion VARCHAR(255) NOT NULL DEFAULT '',
  // image VARCHAR(255) NOT NULL DEFAULT '',

  // status int NOT NULL DEFAULT 0,
  // is_active int NOT NULL DEFAULT 0,
  
  // province_id int NOT NULL DEFAULT 0,
  // city_id int NOT NULL DEFAULT 0,
  // subdistrict_id int NOT NULL DEFAULT 0,
  // ward_id int NOT NULL DEFAULT 0,
  // rt int NOT NULL DEFAULT 0,
  // rw int NOT NULL DEFAULT 0,

  // height int NOT NULL DEFAULT 0,
  // weight int NOT NULL DEFAULT 0,
  // blood_type VARCHAR(255) NOT NULL DEFAULT '',
  // disability int NOT NULL DEFAULT 0,
  // disability_info VARCHAR(255) NULL DEFAULT '',

  // join_date DATE NOT NULL DEFAULT '0000-00-00',
  // details json NULL DEFAULT '{}',
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
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12,
		$13,
		$14,
		$15,
		$16,
		$17,
		$18,
		$19,
		$20,
		$21,
		$22,
		$23,
		$24,
		$25,
		$26,
		$27,
		$28,
		$29
	)`

	
)