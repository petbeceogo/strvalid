package strvalid

import "github.com/asaskevich/govalidator"

func KakaoID(kakaoID string) bool {
	return !govalidator.HasWhitespace(kakaoID)
}
