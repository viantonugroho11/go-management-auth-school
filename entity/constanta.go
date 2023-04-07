package entity

import (
	"database/sql/driver"
	"encoding/json"
	"go-management-auth-school/helper/interfacepkg"
)

type JSONB map[string]interface{}

func NewJSONB(data interface{}) (jsonb JSONB) {
	interfacepkg.Convert(data, &jsonb)
	return
}

func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		// return errors.New("type assertion to []byte failed")
		return nil
	}

	return json.Unmarshal(b, &a)
}
