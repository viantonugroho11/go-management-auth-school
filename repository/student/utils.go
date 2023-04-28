package student

import (
	"context"

	studentRequest "go-management-auth-school/controller/student"
)


func (repo studentRepo) buildingParams(ctx context.Context, parameter *studentRequest.StudentParams) (conditionString string, conditionParam []interface{}) {

	if parameter.ID != "" {
		conditionString += " AND def.id = ?"
		conditionParam = append(conditionParam, parameter.ID)
	}
	if parameter.Nik != "" {
		conditionString += " AND def.nik = ?"
		conditionParam = append(conditionParam, parameter.Nik)
	}
	if parameter.Nisn != "" {
		conditionString += " AND def.nisn = ?"
		conditionParam = append(conditionParam, parameter.Nisn)
	}
	if parameter.Nis != "" {
		conditionString += " AND def.nis = ?"
		conditionParam = append(conditionParam, parameter.Nis)
	}
	if parameter.FirstName != "" {
		conditionString += " AND def.first_name = ?"
		conditionParam = append(conditionParam, parameter.FirstName)
	}
	if parameter.LastName != "" {
		conditionString += " AND def.last_name = ?"
		conditionParam = append(conditionParam, parameter.LastName)
	}
	if parameter.JoinDate != "" {
		conditionString += " AND def.join_date = ?"
		conditionParam = append(conditionParam, parameter.JoinDate)
	}

	return
}