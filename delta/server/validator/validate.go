package validator

import (
	"gopkg.in/go-playground/validator.v9"
)

var (
	validate *validator.Validate
)

func init() {
	validate = newValidator()
}

func newValidator() *validator.Validate {
	var (
		v = validator.New()
	)
	// v.RegisterValidation("minAge", hasMinAge)
	return v
}

// func hasMinAge(fl validator.FieldLevel) bool {
// 	now := time.Now().UTC()
// 	duration := 24 * time.Hour
// 	dob, _ := fl.Field().Interface().(time.Time)
// 	minAge, err := strconv.Atoi(fl.Param())
// 	if err == nil {
// 		minDate := now.AddDate(-minAge, 0, 1).Truncate(duration)
// 		return dob.UTC().Truncate(duration).Before(minDate)
// 	}

// 	return false
// }

// Struct validates a struct
func Struct(s interface{}) error {
	return validate.Struct(s)
}
