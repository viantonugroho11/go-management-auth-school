package database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

func ReplaceSQL(stmt, pattern string, len int) string {
	pattern += ","
	stmt = fmt.Sprintf(stmt, strings.Repeat(pattern, len))
	n := 0
	for strings.IndexByte(stmt, '?') != -1 {
		n++
		param := "$" + strconv.Itoa(n)
		stmt = strings.Replace(stmt, "?", param, 1)
	}
	return strings.TrimSuffix(stmt, ",")
}

func SubstitutePlaceholder(data string, startInt int) (res string) {
	placeholderCount := strings.Count(data, "?")
	res = data
	for i := startInt; i < startInt+placeholderCount; i++ {
		res = strings.Replace(res, "?", "$"+strconv.Itoa(i), 1)
	}
	return res
}

func BuildStrParams(len int, separator string) string {
	var res string
	for i := 0; i < len; i++ {
		if i != 0 {
			res += separator
		}
		res += "?"
	}
	return res
}

func AppendStrArgs(args []string, currArgs *[]interface{}) {
	for _, v := range args {
		*currArgs = append(*currArgs, v)
	}
}

func AppendIntArgs(args []int, currArgs *[]interface{}) {
	for _, v := range args {
		*currArgs = append(*currArgs, v)
	}
}


// how to make utils update
// values := map[string]interface{}{
//     "name": "John",
//     "age":  30,
// }
// updatedData, err := updateDatabase(ctx,tx, db, "users", values, "id = $1",false)
func UpdateDatabase(ctx context.Context,tx *sqlx.Tx, db *sqlx.DB, table string, values map[string]interface{}, condition string, returning bool) (map[string]interface{}, error) {
    query := fmt.Sprintf("UPDATE %s SET ", table)

    var args []interface{}
		var row *sql.Row
		var res sql.Result
		var err error
    i := 1
    for k, v := range values {
        query += fmt.Sprintf("%s = $%d, ", k, i)
        args = append(args, v)
        i++
    }

    query = query[:len(query)-2] // Remove trailing comma and space

    if condition != "" {
        query += fmt.Sprintf("WHERE %s", condition)
    }

		
		if returning == true {
			// Add RETURNING clause to query to retrieve updated data
    	query += " RETURNING *"		
    	// Execute query and retrieve updated data
			if tx != nil {
				row = tx.QueryRowContext(ctx, query, args...)
			}else{
				row = db.QueryRowContext(ctx, query, args...)
			}
    	updatedData := make(map[string]interface{})
    	err = row.Scan(updatedData)
    	if err != nil {
    	    return nil, err
    	}

    	return updatedData, nil
		}
		
		if tx != nil {
			res, err = tx.ExecContext(ctx, query, args...)
		} else {
			res, err = db.ExecContext(ctx, query, args...)
		}
		if err != nil {
			return nil,err
		}
		count, err := res.RowsAffected()
		if err != nil {
			return nil,err
		}
		if count < 1 {
			return nil,sql.ErrNoRows
		}

		return nil,nil  
}
