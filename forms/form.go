package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

// 邮箱验证正则表达式
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

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
			f.Errors.Append(field, "此字段不能为空")
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
		f.Errors.Append(field, fmt.Sprintf("字段长度必须小于%d", max+1))
	}
}

// MinLength 检查字段最小长度，如果小于最小长度，则添加该字段的错误消息
func (f *Form) MinLength(field string, min int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < min {
		f.Errors.Append(field, fmt.Sprintf("字段长度必须大于%d", min-1))
	}
}

func (f *Form) MatchesPattern(pattern *regexp.Regexp, field, msg string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		f.Errors.Append(field, msg)
	}
}
