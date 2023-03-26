package student

import (
	"context"

	studentRequset "go-management-auth-school/controller/student"
	studentEntity "go-management-auth-school/entity/student"

	"go-management-auth-school/helper/database"

	"github.com/jmoiron/sqlx"
)


type studentRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewStudentRepo(dbMaster ,dbSlave *sqlx.DB) *studentRepo {
	return &studentRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo studentRepo) buildingParams(ctx context.Context, parameter *studentRequset.StudentParams) (conditionString string, conditionParam []interface{}) {
	
	if parameter.ID != "" {
		conditionString += " AND id = ?"
		conditionParam = append(conditionParam, parameter.ID)
	}
	if parameter.Nik != "" {
		conditionString += " AND nik = ?"
		conditionParam = append(conditionParam, parameter.Nik)
	}
	if parameter.Nisn != "" {
		conditionString += " AND nisn = ?"
		conditionParam = append(conditionParam, parameter.Nisn)
	}
	if parameter.Nis != "" {
		conditionString += " AND nis = ?"
		conditionParam = append(conditionParam, parameter.Nis)
	}
	if parameter.FirstName != "" {
		conditionString += " AND first_name = ?"
		conditionParam = append(conditionParam, parameter.FirstName)
	}
	if parameter.LastName != "" {
		conditionString += " AND last_name = ?"
		conditionParam = append(conditionParam, parameter.LastName)
	}
	if parameter.JoinDate != "" {
		conditionString += " AND join_date = ?"
		conditionParam = append(conditionParam, parameter.JoinDate)
	}

	return 
}


func (repo studentRepo) SelectAll(ctx context.Context, parameter *studentRequset.StudentParams) (data []studentEntity.Student, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := studentEntity.SelectUser + ` WHERE def.delete_at is null ` + whereStatment + 
	` ORDER BY ` + parameter.OrderBy + ` ` + parameter.Sort + `, def.id ` + parameter.Sort

	query = database.SubstitutePlaceholder(query, 1)
	rows, err := repo.DbSlave.QueryContext(ctx, query, conditionParam...)
	if err != nil {
		return
	}

	for rows.Next() {
		temp := studentEntity.Student{}
		err = temp.ScanRows(rows,nil)
		if err != nil {
			return
		}
		data = append(data, temp)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}
