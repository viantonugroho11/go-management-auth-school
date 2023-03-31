package mappingstudent

import (
)

type MappingStudent struct {
	ID int `json:"id"`
	Indentity string `json:"indentity"`
	ClassID string `json:"class_id"`
	Type string `json:"type"`
}