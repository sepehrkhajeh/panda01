package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// CustomValidator ساختار اصلی برای ولیدیشن سفارشی
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate متد اصلی ولیدیشن برای Echo
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// CustomValidation تابعی برای بررسی فرمت خاص رشته‌ها
func PubKeyValidation(fl validator.FieldLevel) bool {
	// ریجکس برای رشته‌ای که فقط شامل اعداد 0-9 و حروف a-f (کوچک) باشد و دقیقاً ۶۴ کاراکتر طول داشته باشد
	regex := regexp.MustCompile(`^[a-f0-9]{64}$`)
	return regex.MatchString(fl.Field().String())
}

// NewValidator ایجاد و مقداردهی اولیه Validator
func NewValidator() *validator.Validate {
	v := validator.New()
	// ثبت ولیدیشن سفارشی
	v.RegisterValidation("customString", PubKeyValidation)

	return v
}
