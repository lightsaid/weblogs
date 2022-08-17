package validator

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// 表单数据验证

// Form 一个 FormValidator 结构，包含 url.Values 和 errors, 用于验证表单数据
type FormValidator struct {
	url.Values
	Errors errors
}

// New 实例化一个 FormValidator 结构
func NewFormValidator(data url.Values) *FormValidator {
	return &FormValidator{
		data,
		errors(map[string][]string{}),
	}
}

// Valid 验证是否存在错误，存在：false， 反之：true
func (f *FormValidator) Valid() bool {
	return len(f.Errors) == 0
}

// Required 非空校验
func (f *FormValidator) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, fmt.Sprintf("%s字段不能为空", field))
		}
	}
}

// MinLength 检查字段最小长度
func (f *FormValidator) MinLength(field string, length int) bool {
	fd := f.Get(field)
	if len(fd) < length {
		f.Errors.Add(field, fmt.Sprintf("%s字段长度必须大于%d", field, length))
		return false
	}
	return true
}

// MaxLength 检查字段最大长度
func (f *FormValidator) MaxLength(field string, length int) bool {
	fd := f.Get(field)
	if len(fd) > length {
		f.Errors.Add(field, fmt.Sprintf("%s字段长度必须小于%d", field, length))
		return false
	}
	return true
}

// IsEmail 检查email是否有效, 如果无效则添加错误信息
func (f *FormValidator) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "邮箱地址无效")
	}
}
