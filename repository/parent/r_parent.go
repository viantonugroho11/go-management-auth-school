package parent

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	parentController "go-management-auth-school/controller/parent"
	parentEntity "go-management-auth-school/entity/parent"
)

type parentRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewParentRepo(dbMaster, dbSlave *sqlx.DB) *parentRepo {
	return &parentRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo parentRepo) buildingParams(ctx context.Context, parameter *parentController.ParentParams) (conditionString string, conditionParam []interface{}) {

	if parameter.NIK != "" {
		conditionString += " AND def.nik = ?"
		conditionParam = append(conditionParam, parameter.NIK)
	}
	if parameter.FirstName != "" {
		conditionString += " AND def.first_name = ?"
		conditionParam = append(conditionParam, parameter.FirstName)
	}
	if parameter.LastName != "" {
		conditionString += " AND def.last_name = ?"
		conditionParam = append(conditionParam, parameter.LastName)
	}
	if parameter.StudentID != "" {
		conditionString += " AND def.student_id = ?"
		conditionParam = append(conditionParam, parameter.StudentID)
	}

	return
}

func (repo parentRepo) CreateTx(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = repo.DbMaster.BeginTxx(ctx, nil)
	return
}

func (repo parentRepo) FindAll(ctx context.Context, params *parentController.ParentParams) (data []parentEntity.Parent, err error) {
	// build query here
	return
}

func (repo parentRepo) SelectAll(ctx context.Context, parameter *parentController.ParentParams) (data []parentEntity.Parent, err error) {
	// build query here
	return
}

func (repo parentRepo) FindOne(ctx context.Context, params *parentController.ParentParams) (data parentEntity.Parent, err error) {
	// build query here
	return
}

// create parent
func (repo parentRepo) Create(ctx context.Context, tx *sqlx.Tx, params *parentEntity.Parent) (err error) {
	// build query here
	queries := InsertParent
	uuidRandom := uuid.New().String()
	_, err = tx.ExecContext(ctx, queries,
		uuidRandom,
		params.FirstName,
		params.LastName,
		params.Type,
		params.NIK,
		params.Gender,
		params.Phone,
		params.WorkID,
		params.WorkName,
		params.Income,
		params.StudentID,
		params.Image,
	)
	if err != nil {
		return
	}
	return
}
