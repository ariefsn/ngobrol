package validator

import (
	"github.com/ariefsn/ngobrol/entities"
	val "github.com/go-playground/validator/v10"
)

var validate *val.Validate

func InitValidator() {
	validate = val.New(val.WithRequiredStructEnabled())
	validate.RegisterValidation("password", ValidatePassword)
}

func Validator() *val.Validate {
	return validate
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func ValidateVar(field interface{}, tag string) error {
	return validate.Var(field, tag)
}

func ValidateVarMap(data, rules entities.M) entities.M {
	return validate.ValidateMap(data, rules)
}

func ParseValidationError(err error) val.ValidationErrors {
	if err != nil {
		parsed := err.(val.ValidationErrors)
		return parsed
	}

	return nil
}
