package class

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	classController "go-management-auth-school/controller/class"
	classEntity "go-management-auth-school/entity/class"
)

type classRepo struct {
	DbMaster *sqlx.DB
	DbSlave  *sqlx.DB
}

func NewClassRepo(dbMaster, dbSlave *sqlx.DB) *classRepo {
	return &classRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
	}
}

func (repo classRepo) buildingParams(ctx context.Context, parameter *classController.ClassParams) (conditionString string, conditionParam []interface{}) {

	if parameter.ID != "" {
		conditionString += " AND def.id = ?"
		conditionParam = append(conditionParam, parameter.ID)
	}

	return
}

func (repo classRepo) FindAll(ctx context.Context, params *classController.ClassParams) (data []classEntity.Class, err error) {
	// build query here
	return
}

func (repo classRepo) SelectAll(ctx context.Context, parameter *classController.ClassParams) (data []classEntity.Class, err error) {
	// build query here
	whereStatment, conditionParam := repo.buildingParams(ctx, parameter)
	query := classEntity.SelectStatment + ` WHERE def.deleted_at IS NULL ` + whereStatment

	rows, err := repo.DbSlave.QueryContext(ctx, query, conditionParam...)
	if err != sql.ErrNoRows && err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		d := classEntity.Class{}
		err = d.ScanRows(rows, nil)
		if err != sql.ErrNoRows && err != nil {
			return nil, err
		}
		data = append(data, d)
	}

	return
}

func (repo classRepo) FindOne(ctx context.Context, params *classController.ClassParams) (data classEntity.Class, err error) {
	// build query here
	whereStatment, conditionParam := repo.buildingParams(ctx, params)
	query := classEntity.SelectStatment + ` WHERE def.deleted_at IS NULL ` + whereStatment

	row := repo.DbSlave.QueryRowContext(ctx, query, conditionParam...)

	err = data.ScanRows(nil, row)
	if err != sql.ErrNoRows && err != nil {
		return data, err
	}

	return
}

func (repo classRepo) Create(ctx context.Context, tx *sqlx.Tx, params *classEntity.Class) (err error) {
	// build query here
	queries := InsertClass
	_, err = tx.ExecContext(ctx, queries, params.Name, params.MajorID, params.Level)
	if err != nil {
		return err
	}
	return
}

func (repo classRepo) CreateTx(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = repo.DbMaster.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return tx, err
	}
	return
}

// delete
func (repo classRepo) Delete(ctx context.Context, tx *sqlx.Tx, params *classEntity.Class) (err error) {
	// build query here
	queries := `UPDATE ` + classEntity.Table + ` SET deleted_at = NOW() WHERE id = ?`
	if tx == nil {
		_, err = repo.DbMaster.ExecContext(ctx, queries, params.ID)
		if err != nil {
			return err
		}
		return
	}
	_, err = tx.ExecContext(ctx, queries, params.ID)
	if err != nil {
		return err
	}
	return
}

// update
func (repo classRepo) Update(ctx context.Context, tx *sqlx.Tx, params *classEntity.Class) (err error) {
	// build query here
	queries := `UPDATE ` + classEntity.Table + ` SET name = ?, major_id = ?, level = ? WHERE id = ?`
	if tx == nil {
		_, err = repo.DbMaster.ExecContext(ctx, queries, params.Name, params.MajorID, params.Level, params.ID)
		if err != nil {
			return err
		}
		return
	}
	_, err = tx.ExecContext(ctx, queries, params.Name, params.MajorID, params.Level, params.ID)
	if err != nil {
		return err
	}
	return
}