package user

import (
	// "context"

	"context"
	"database/sql"
	userRequset "go-management-auth-school/controller/user"
	userEntity "go-management-auth-school/entity/user"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"go-management-auth-school/helper/database"
)

type userRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewStudentRepo(dbMaster ,dbSlave *sqlx.DB) *userRepo {
	return &userRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo userRepo) buildingParams(ctx context.Context, parameter *userRequset.UserParams) (conditionString string, conditionParam []interface{}) {
	
	if parameter.ID != 0 {
		conditionString += " AND id = ?"
		conditionParam = append(conditionParam, parameter.ID)
	}
	if parameter.Username != "" {
		conditionString += " AND username = ?"
		conditionParam = append(conditionParam, parameter.Username)
	}
	if parameter.IdentityID != "" {
		conditionString += " AND identity_id = ?"
		conditionParam = append(conditionParam, parameter.IdentityID)
	}
	if parameter.Permission != "" {
		conditionString += " AND permission = ?"
		conditionParam = append(conditionParam, parameter.Permission)
	}

	return 
}

func (repo userRepo) CreateTx(ctx context.Context) (tx *sqlx.Tx, err error) {
	return repo.DbMaster.BeginTxx(ctx, &sql.TxOptions{})
}

func (repo userRepo) SelectAll(ctx context.Context, parameter *userRequset.UserParams) (data []userEntity.User, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := userEntity.SelectUser + ` WHERE def.delete_at is null ` + whereStatment + 
	` ORDER BY ` + parameter.OrderBy + ` ` + parameter.Sort + `, def.id ` + parameter.Sort

	query = database.SubstitutePlaceholder(query, 1)
	rows, err := repo.DbSlave.QueryContext(ctx, query, conditionParam...)
	if err != nil {
		return
	}

	for rows.Next() {
		temp := userEntity.User{}
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

func(repo userRepo) FindOne(ctx context.Context, parameter *userRequset.UserParams) (data userEntity.User, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := userEntity.SelectUser + ` WHERE def.delete_at is null ` + whereStatment + 
	` ORDER BY ` + parameter.OrderBy + ` ` + parameter.Sort + `, def.id ` + parameter.Sort

	query = database.SubstitutePlaceholder(query, 1)
	row := repo.DbSlave.QueryRowContext(ctx, query, conditionParam...)
	if err != nil {
		return
	}
	err = data.ScanRows(nil,row)
	if err != nil {
		return
	}

	return
}

func(repo userRepo) Create(ctx context.Context,tx *sqlx.Tx, input *userEntity.User) (res string, err error) {
	uuid := uuid.New()
	query := `INSERT INTO `+userEntity.Table+` (id, username, password, identity_id, permission) VALUES ($1,$2,$3,$4) returning identity_id`
	err = tx.QueryRowContext(ctx, query, uuid, input.Username, input.Password, input.IdentityID, input.Permission).Scan(&res)
	if err != nil {
		return
	}

	return
}