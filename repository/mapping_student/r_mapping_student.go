package mapping_student

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	mapStudentController "go-management-auth-school/controller/mapping_student"
	mapStudentEntity "go-management-auth-school/entity/mapping_student"
)

type mpStudentRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewMpStudentRepo(dbMaster, dbSlave *sqlx.DB) *mpStudentRepo {
	return &mpStudentRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo mpStudentRepo) buildingParams(ctx context.Context, parameter *mapStudentController.MappingStudentParams) (conditionString string, conditionParam []interface{}) {

	if parameter.ID != "" {
		conditionString += " AND def.id = ?"
		conditionParam = append(conditionParam, parameter.ID)
	}
	if parameter.Identity != "" {
		conditionString += " AND def.identity = ?"
		conditionParam = append(conditionParam, parameter.Identity)
	}
	if parameter.ClassID != "" {
		conditionString += " AND def.class_id = ?"
		conditionParam = append(conditionParam, parameter.ClassID)
	}
	if parameter.Type != "" {
		conditionString += " AND def.type = ?"
		conditionParam = append(conditionParam, parameter.Type)
	}

	return
}

// create tx
func (repo mpStudentRepo) CreateTx(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = repo.DbMaster.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}
	return
}

func (repo mpStudentRepo) FindAll(ctx context.Context, params *mapStudentController.MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error) {
	// build query here
	return
}

func (repo mpStudentRepo) SelectAll(ctx context.Context, parameter *mapStudentController.MappingStudentParams) (data []mapStudentEntity.MappingStudent, err error) {
	// build query here
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := mapStudentEntity.SelectMapStudent + ` WHERE def.deleted_at IS NULL ` + whereStatment + ` ` + mapStudentEntity.GroupStatement +
		` ORDER BY def.id` + parameter.OrderBy

	// query = database.SubstitutePlaceholder(query, 1)
	rows, err := repo.DbSlave.QueryContext(ctx, query, conditionParam...)
	if err != sql.ErrNoRows {
		return
	}
	defer rows.Close()

	for rows.Next() {
		temp := mapStudentEntity.MappingStudent{}
		err = temp.ScanRows(rows, nil)
		if err != sql.ErrNoRows {
			return
		}
		data = append(data, temp)
	}

	err = rows.Err()
	if err != sql.ErrNoRows {
		return
	}

	return
}

func (repo mpStudentRepo) FindOne(ctx context.Context, parameter *mapStudentController.MappingStudentParams) (data mapStudentEntity.MappingStudent, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := mapStudentEntity.SelectMapStudent + ` WHERE def.deleted_at IS NULL ` + whereStatment

	// query = database.SubstitutePlaceholder(query, 1)
	row := repo.DbSlave.QueryRowContext(ctx, query, conditionParam...)
	err = data.ScanRows(nil, row)
	if err != sql.ErrNoRows {
		return
	}

	return
}

func (repo mpStudentRepo) Create(ctx context.Context, tx *sqlx.Tx, params *mapStudentEntity.MappingStudentReq) (err error) {
	// build query here
	queries := InsertMapStudent
	uuidRandom := uuid.New().String()
	_, err = tx.ExecContext(ctx, queries, uuidRandom, params.IdentityID, params.ClassID, params.Type)
	if err != nil {
		return
	}
	return
}
