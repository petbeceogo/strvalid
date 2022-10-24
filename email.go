package validator

import "github.com/asaskevich/govalidator"

func Email(email string) bool {
	return govalidator.IsEmail(email)
}
