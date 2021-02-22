package utils

import (
	"regexp"

	"github.com/gookit/validate"
)

func ValidatePassword(ps string) bool {
	if len(ps) < 8 || len(ps) > 21 {
		return false
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+-=|_.]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return false
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		return false
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		return false
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return false
	}
	return true
}

// Repassword 确认密码
//	@see validate:"required|repassword:Password" message:"required:确认密码不能为空|repassword:两次输入的密码不一致"
func Repassword(valid *validate.Validation) interface{} {
	return func(value, attribute string) bool {
		if password, has := valid.Get(attribute); has {
			if pwd, ok := password.(string); ok {
				return pwd == value
			} else {
				return false
			}
		} else {
			return false
		}
		return true
	}
}

func init() {
	validate.AddValidator("password", ValidatePassword)
}
