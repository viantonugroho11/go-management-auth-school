package str

import (
	"unicode"
)

func ValidatePassword(pass string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)
 
	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}
 
	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}
 
	return true
}