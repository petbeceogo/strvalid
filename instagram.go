package validator

import "github.com/asaskevich/govalidator"

func InstagramID(instaID string) bool {
	return !govalidator.HasWhitespace(instaID)
}
