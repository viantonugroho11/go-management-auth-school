package user

import (
	// "context"

	"context"
	"database/sql"
	userRequset "go-management-auth-school/controller/user"
	userEntity "go-management-auth-school/entity/user"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	// "go-management-auth-school/helper/database"
)

type userRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewUserRepo(dbMaster ,dbSlave *sqlx.DB) *userRepo {
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
	query := userEntity.SelectUser + ` WHERE def.deleted_at is null ` + whereStatment + 
	` ORDER BY def.id` + parameter.OrderBy 

	// query = database.SubstitutePlaceholder(query, 1)
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
	query := userEntity.SelectUser + ` WHERE def.deleted_at IS NULL` + whereStatment 

	// query = database.SubstitutePlaceholder(query, 1)
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
	uuidRandom := uuid.New().String()
	query := `INSERT INTO `+userEntity.Table+` (id, username, password, identity_id, permission, status) VALUES (?,?,?,?,?,?)`
	_, err = tx.ExecContext(ctx, query, uuidRandom, input.Username, input.Password, input.IdentityID, input.Permission, 1)
	if err != nil {
		return
	}
	res = input.IdentityID

	return
}

func (repo userRepo) UpdateUsername(ctx context.Context,tx *sqlx.Tx, input *userEntity.User) (err error) {
	query := `UPDATE `+userEntity.Table+` SET username = ? WHERE identity_id = ?`
	_, err = tx.ExecContext(ctx, query, input.Username, input.IdentityID)
	if err != nil {
		return
	}

	return
}

func (repo userRepo) UpdatePassword(ctx context.Context,tx *sqlx.Tx, input *userEntity.User) (err error) {
	query := `UPDATE `+userEntity.Table+` SET password = ? WHERE identity_id = ?`
	_, err = tx.ExecContext(ctx, query, input.Password, input.IdentityID)
	if err != nil {
		return
	}

	return
}

func (repo userRepo) UpdatePermission(ctx context.Context,tx *sqlx.Tx, input *userEntity.User) (err error) {
	query := `UPDATE `+userEntity.Table+` SET permission = ? WHERE identity_id = ?`
	_, err = tx.ExecContext(ctx, query, input.Permission, input.IdentityID)
	if err != nil {
		return
	}

	return
}

//last login
func (repo userRepo) UpdateLastLogin(ctx context.Context,tx *sqlx.Tx, input *userEntity.User) (err error) {
	//time now local jakarta
	timeNow := time.Now().In(time.FixedZone("Asia/Jakarta", 7*60*60)).Format("2006-01-02 15:04:05")
	query := `UPDATE `+userEntity.Table+` SET last_login = ? WHERE identity_id = ?`
	_, err = tx.ExecContext(ctx, query, timeNow, input.IdentityID)
	if err != nil {
		return
	}

	return
}