package parent

import (
	"context"
	"database/sql"

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
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := parentEntity.SelectParent + ` WHERE def.deleted_at IS NULL ` + whereStatment + ` ` + parentEntity.GroupStatement +
		` ORDER BY def.id` + parameter.OrderBy

	// query = database.SubstitutePlaceholder(query, 1)
	rows, err := repo.DbSlave.QueryContext(ctx, query, conditionParam...)
	if err != sql.ErrNoRows && err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		temp := parentEntity.Parent{}
		err = temp.ScanRows(rows, nil)
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

func (repo parentRepo) FindOne(ctx context.Context, params *parentController.ParentParams) (data parentEntity.Parent, err error) {
	whereStatment, conditionParam := repo.buildingParams(ctx, params)
	query := parentEntity.SelectParent + ` WHERE def.deleted_at IS NULL ` + whereStatment

	// query = database.SubstitutePlaceholder(query, 1)
	row := repo.DbSlave.QueryRowContext(ctx, query, conditionParam...)
	err = data.ScanRows(nil, row)
	if err != sql.ErrNoRows && err != nil {
		return
	}

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

// create bulk parent
func (repo parentRepo) CreateBulk(ctx context.Context, tx *sqlx.Tx, parents []*parentEntity.Parent) (err error) {
	// Prepare the query
	queries := `INSERT INTO ` + parentEntity.TableName + ` (
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
	) VALUES `
	
	// Prepare the values slice
	values := make([]interface{}, 0, len(parents)*12) // Assuming each parent has 12 fields
	
	// Build the query placeholders and collect the values
	for _, parent := range parents {
		uuidRandom := uuid.New().String()
		queries += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"
		values = append(values,
			uuidRandom,
			parent.FirstName,
			parent.LastName,
			parent.Type,
			parent.NIK,
			parent.Gender,
			parent.Phone,
			parent.WorkID,
			parent.WorkName,
			parent.Income,
			parent.StudentID,
			parent.Image,
		)
	}
	
	// Remove the trailing comma
	queries = queries[:len(queries)-1]
	
	// Execute the bulk insert query
	_, err = tx.ExecContext(ctx, queries, values...)
	if err != nil {
		return 
	}
	return 
}

// update parent
func (repo parentRepo) Update(ctx context.Context, tx *sqlx.Tx, params *parentEntity.Parent) (err error) {
	query := `UPDATE ` + parentEntity.TableName + ` SET
		first_name = ?,
		last_name = ?,
		type = ?,
		nik = ?,
		gender = ?,
		phone = ?,
		work_id = ?,
		work_name = ?,
		income = ?,
		student_id = ?,
		image = ?
		WHERE id = ?`
	_, err = tx.ExecContext(ctx, query,
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
		params.ID,
	)
	if err != nil {
		return
	}
	return
}

// delete parent
func (repo parentRepo) Delete(ctx context.Context, tx *sqlx.Tx, params *parentEntity.Parent) (err error) {
	query := `UPDATE ` + parentEntity.TableName + ` SET deleted_at = now() WHERE id = ?`
	_, err = tx.ExecContext(ctx, query, params.ID)
	if err != nil {
		return
	}
	return
}


