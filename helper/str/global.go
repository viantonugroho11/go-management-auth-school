package str

import (
	"fmt"
	timeHelper "go-management-auth-school/helper/time"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	LoanRepayment = "loan_repayment"
	LoanPartial   = "loan_partial"
	Deposit       = "deposit"
	Withdrawal    = "withdrawal"
)

// ShowString ...
func ShowString(isShow bool, data string) string {
	if isShow {
		return data
	}

	return ""
}

// EmptyString ...
func EmptyString(text string) *string {
	if text == "" {
		return nil
	}
	return &text
}

// EmptyInt ...
func EmptyInt(number int) *int {
	if number == 0 {
		return nil
	}
	return &number
}

// EmptyFloat ...
func EmptyFloat(amount float64) *float64 {
	if amount == 0 {
		return nil
	}
	return &amount
}

// StringToInt ...
func StringToInt(data string) int {
	res, err := strconv.Atoi(data)
	if err != nil {
		res = 0
	}

	return res
}

// StringToBool ...
func StringToBool(data string) bool {
	res, _ := strconv.ParseBool(data)
	return res
}

// StringToBoolString ...
func StringToBoolString(data string) string {
	res, _ := strconv.ParseBool(data)
	if res {
		return "true"
	}

	return "false"
}

// RandomString ...
func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()

	return str
}

// IsActive ...
func IsActive(data string) *string {

	var isActive string
	res, err := strconv.ParseBool(data)

	if err != nil {
		isActive = ""
		return &isActive
	}

	if res {
		isActive = "and is_active = 'true'"
	} else {
		isActive = "and is_active = 'false'"
	}

	return &isActive
}

// Unique ...
func Unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

// UniqueInt ...
func UniqueInt(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

// CheckEmail ...
func CheckEmail(text string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(text)
}

// IsValidUUID ...
func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func IntNilChecker(input *int) int {
	if input == nil {
		return 0
	}
	return *input
}

func FloatNilChecker(input *float64) float64 {
	if input == nil {
		return 0
	}
	return *input
}

func StringNilChecker(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}

func JoinWithQuotes(data []string, quotes, separator string) (res string) {
	for i, v := range data {
		if i != 0 {
			res += separator
		}

		res += quotes + v + quotes
	}

	return
}

func CheckNumber(text string) bool {
	r := regexp.MustCompile(`^[0-9]+$`)
	return r.MatchString(text)
}

func CheckBulkInt(data []string) bool {
	for _, v := range data {
		if _, err := strconv.Atoi(v); err != nil {
			return false
		}
	}

	return true
}

func ConvertStringPointer(word *string) string {
	if word == nil {
		return ""
	}
	return *word
}

func NormalizePhone(text string) string {
	re := regexp.MustCompile(`[/\D/g]`)
	text = re.ReplaceAllString(text, "")
	if len(text) < 8 {
		return ""
	}
	if text[:2] == "62" {
		text = strings.Replace(text, "62", "0", 1)
	}
	return text
}

// MultipleValueParameter
func MultipleValueParameter(r *http.Request, param string) []string {
	params := r.FormValue(param)
	if params == "" { // avoid strings.Split return []string{""} instead []string{}
		return []string{}
	}

	return strings.Split(params, ",")
}

// ArrayStringToInt
func ArrayStringToInt(arr []string) []int {
	var arrInt = []int{}

	for _, i := range arr {
		j, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		arrInt = append(arrInt, j)
	}
	return arrInt
}

// GenerateTransactionCode
func GenerateTransactionCode(types string, count int) (code string) {
	currDate, _ := timeHelper.In(time.Now(), timeHelper.JAKARTA_TIME_LOCATION)
	date := currDate.Format("060102")
	return fmt.Sprintf("AO"+types+date+"%06d", count)
}

// ConvertFloat64Pointer
func ConvertFloat64Pointer(number *float64) float64 {
	if number == nil {
		return 0
	}
	return *number
}

// RemoveString
func RemoveString(s string, index, length int) string {
	return s[:index] + s[index+length:]
}
