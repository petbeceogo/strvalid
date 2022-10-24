package validator

import "github.com/asaskevich/govalidator"

func Phone(phone string) bool {
	return govalidator.StringMatches(phone, "^01([0|1|6|7|8|9])-?([0-9]{3,4})-?([0-9]{4})$")
}
