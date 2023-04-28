package major

import (
	majorEntity "go-management-auth-school/entity/major"
)

var (
	InsertMajor = `INSERT INTO ` + majorEntity.Table + ` (name) VALUES (?)`
)
