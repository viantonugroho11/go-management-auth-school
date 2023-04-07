package sqlscan

import (
	"encoding/json"
)

// this model is using to receive json value form database
// and make it to be struct that we wanted
// this can be easily using on (row / rows).Scan
//row.Scan(
// 	&sqlscan.UnmarshalData{to: &res.Any}
// )
type UnmarshalData struct {
	To interface{} // can be anything
}

// implement sql.Scanner interface
func (a *UnmarshalData) Scan(value interface{}) error {
	b, ok := value.([]byte) // postgres return jsonb type as bytes
	if !ok {
		// return errors.New("type assertion to []byte failed")
		return nil
	}

	return json.Unmarshal(b, &a.To)
}
