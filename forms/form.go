package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// Form 自定义Form结构，匿名嵌套 url.Values，和存放验证不通过的Errors
type Form struct {
	url.Values
	Errors errors
}

// New 实例化，参数data: 例如POST请求中，data是r.PostForm，将请求的参数存放到自定义Form里面
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Valid 验证 Form Errors 是否存在错误，没有错误返回 true
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Required 验证哪些字段为必填
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "此字段不能为空")
		}
	}
}

// MaxLength 检查字段最大长度，如果超出最大长度，则添加该字段的错误消息
func (f *Form) MaxLength(field string, max int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > max {
		f.Errors.Add(field, fmt.Sprintf("字段长度不允许超出: %d", max))
	}
}

// MinLength 检查字段最小长度，如果小于最小长度，则添加该字段的错误消息
func (f *Form) MinLength(field string, min int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < min {
		f.Errors.Add(field, fmt.Sprintf("字段长度必须大于等于: %d", min))
	}
}
