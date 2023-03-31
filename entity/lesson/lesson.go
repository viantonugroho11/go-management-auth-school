package lesson

import (
)

type Lesson struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}