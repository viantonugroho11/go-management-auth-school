package verify_token

import (
	"context"
	verifyTokenRequest "go-management-auth-school/controller/verify_token"
	verifyTokenEntity "go-management-auth-school/entity/verify_token"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type verifyTokenRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewVerifyTokenRepo(dbMaster *sqlx.DB, dbSlave *sqlx.DB) *verifyTokenRepo {
	return &verifyTokenRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo verifyTokenRepo) buildingParams(ctx context.Context, parameter *verifyTokenRequest.VerifyTokenParams) (conditionString string, conditionParam []interface{}) {

	if parameter.Token != "" {
		conditionString += " AND def.token = ?"
		conditionParam = append(conditionParam, parameter.Token)
	}
	if parameter.Identity != "" {
		conditionString += " AND def.identity_id = ?"
		conditionParam = append(conditionParam, parameter.Identity)
	}
	return
}
func (repo verifyTokenRepo) FindOne(ctx context.Context, parameter *verifyTokenRequest.VerifyTokenParams) (data verifyTokenEntity.VerifyToken, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := verifyTokenEntity.SelectStatement + ` WHERE def.deleted_at is null ` + whereStatment +
		` ORDER BY def.id` + parameter.OrderBy
	// query = database.SubstitutePlaceholder(query, 1)
	row := repo.DbSlave.QueryRowContext(ctx, query, conditionParam...)

	err = data.ScanRows(nil, row)
	if err != nil {
		return
	}
	return
}

// create TX
func (repo verifyTokenRepo) CreateTx(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = repo.DbMaster.BeginTxx(ctx, nil)
	if err != nil {
		return
	}
	return
}

func (repo verifyTokenRepo) Create(ctx context.Context, tx *sqlx.Tx, parameter *verifyTokenEntity.VerifyToken) (err error) {
	queries := InsertVerifyToken
	uuidRandom := uuid.New().String()
	if tx != nil {
		_, err = tx.QueryContext(ctx, queries, uuidRandom, parameter.Identity, parameter.Token, parameter.ExpiredAt)
		if err != nil {
			return err
		}
		return
	}
	_, err = repo.DbMaster.QueryContext(ctx, queries, uuidRandom, parameter.Identity, parameter.Token, parameter.ExpiredAt)
	if err != nil {
		return err
	}
	return

}

func (repo verifyTokenRepo) Delete(ctx context.Context, tx *sqlx.Tx, parameter *verifyTokenEntity.VerifyToken) (err error) {
	queries := DeleteVerifyToken
	if tx != nil {
		_, err = tx.QueryContext(ctx, queries, parameter.Identity)
		if err != nil {
			return err
		}
		return
	}
	_, err = repo.DbMaster.QueryContext(ctx, queries, parameter.Identity)
	if err != nil {
		return err
	}
	return
}
