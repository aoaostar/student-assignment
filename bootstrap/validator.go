package bootstrap

import (
	"StudentAssignment/pkg/validator"
	v "github.com/go-playground/validator/v10"
)

func InitValidator()  {


	validator.Validate = v.New()
}
