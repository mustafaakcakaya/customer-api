package utils

import (
	"CustomerAPI/internal/customer/types"
	"CustomerAPI/pkg/errors"
	"fmt"
	"regexp"
)

func ValidateRequest(ctx echo.Context, request interface{}) (result interface{}, err error) {

	err = ctx.Bind(request)
	if err != nil {
		if _, ok := err.(*echo.HTTPError); ok {
			return nil, errors.ValidatorError.WrapDesc(err.(*echo.HTTPError).Message.(string))
		}
		return errors.UnknownError.Wrap(err), nil
	}

	v := validator.New()

	v.RegisterStructValidation(customerValidations, types.CustomerUpsertRequest{})

	if err = v.Struct(request); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			desc := ""
			for _, err := range err.(validator.ValidationErrors) {
				desc = fmt.Sprintf("Problem is: '%s' , Hint: '%s' ", err.Field(), err.ActualTag())
			}
			return nil, errors.ValidatorError.WrapDesc(desc)
		}
		return errors.UnknownError.Wrap(err), nil
	}

	return request, nil
}

func customerValidations(sl validator.StructLevel) {
	model := sl.Current().Interface().(types.CustomerUpsertRequest)

	if !isEmailValid(model.Email) {
		panic(errors.ValidatorError.WrapDesc("Email is not valid!"))
	}
	if *model.Address.CityCode > 81 || *model.Address.CityCode < 1 {
		panic(errors.ValidatorError.WrapDesc("CityCode should be between 1 and 81!"))
	}
}
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
