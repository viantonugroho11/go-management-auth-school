package str

import (
	"log"
	"strconv"
)

func ToFloat64(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println("str.ToFloat64", err)
	}
	return res
}
