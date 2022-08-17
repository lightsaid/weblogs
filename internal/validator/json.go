package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/asaskevich/govalidator"
)

// JsonValidator 验证的JSON数据的结构体, 仅仅对一层JOSN验证
// 有效：{"age": 10, "name":"zhangsan", "ok": true}
// 无效：{"list":[1,2,3]}、{"person":{"age":100, "name":"zhangsan"}}
type JsonValidator struct {
	Data   map[string]interface{}
	Errors errors
}

// NewJsonValidator 实例化一个JSON验证器，参数 data 必须是struct或者是struct指针值
// 将传递过来的data转化为 map[string]interface{} 存起来.
// 如果参数不满足要求，如传一个nil, 始终会返回一个初始化的JsonValidator实例，即使err不等于nil
func NewJsonValidator(data interface{}) (*JsonValidator, error) {
	object := make(map[string]interface{})
	jvalid := &JsonValidator{
		Data:   object,
		Errors: make(map[string][]string),
	}

	if data == nil {
		return jvalid, fmt.Errorf("data Kind expect got reflect.Struct or reflect.Ptr, but %v", nil)
	}

	dataType := reflect.TypeOf(data)
	dataVal := reflect.ValueOf(data)

	switch dataType.Kind() {
	case reflect.Struct:
	case reflect.Ptr:
		dataType = dataType.Elem()
		dataVal = dataVal.Elem()
		if dataType.Kind() != reflect.Struct {
			// fmt.Println(dataType.Kind())
			return jvalid, fmt.Errorf("data Kind expect got reflect.Struct or reflect.Ptr, but %s", dataType.Kind())
		}
	default:
		return jvalid, fmt.Errorf("data Kind expect got reflect.Struct or reflect.Ptr, but %s", dataType.Kind())
	}

	for i := 0; i < dataType.NumField(); i++ {
		fieldName := dataType.Field(i).Name
		fieldVal := dataVal.Field(i)
		tag := dataType.Field(i).Tag.Get("json")
		// 没有 json tag， 使用 Struct 字段名
		if len(tag) == 0 {
			tag = fieldName
		}
		switch fieldVal.Kind() {
		case reflect.String:
			fieldValue := fieldVal.String()
			object[tag] = fieldValue
		case reflect.Bool:
			fieldValue := fieldVal.Bool()
			object[tag] = fieldValue
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			fieldValue := fieldVal.Int()
			object[tag] = fieldValue
		case reflect.Float32, reflect.Float64:
			fieldValue := fieldVal.Float()
			object[tag] = fieldValue
		default:
			object[tag] = fieldVal.String()
		}
	}
	jvalid.Data = object
	return jvalid, nil
}

// Valid 验证是否存在错误，存在：false， 反之：true
func (v *JsonValidator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError 添加错误信息
func (v *JsonValidator) AddError(field, message string) {
	v.Errors.Add(field, message)
}

// Check 检查是否错误，如果是错误则添加错误信息
// 例如：Check(age >= 18, "age", "禁止未成年人入内")
func (v *JsonValidator) Check(exp bool, field, message string) {
	if !exp {
		v.AddError(field, message)
	}
}

// Required 非空校验
func (v *JsonValidator) Required(fields ...string) {
	for _, field := range fields {
		value := v.Data[field]
		if strings.TrimSpace(fmt.Sprint(value)) == "" {
			v.AddError(field, fmt.Sprintf("%s字段不能为空", field))
		}
	}
}

// MinLength 检查字段最小长度
func (v *JsonValidator) MinLength(field string, length int) bool {
	value := v.Data[field]
	if len(fmt.Sprint(value)) < length {
		v.AddError(field, fmt.Sprintf("%s字段长度必须大于%d", field, length))
		return false
	}
	return true
}

// MaxLength 检查字段最大长度
func (v *JsonValidator) MaxLength(field string, length int) bool {
	value := v.Data[field]
	if len(fmt.Sprint(value)) > length {
		v.AddError(field, fmt.Sprintf("%s字段长度必须小于%d", field, length))
		return false
	}
	return true
}

// IsEmail 检查email是否有效, 如果无效则添加错误信息
func (v *JsonValidator) IsEmail(field string) {
	value := fmt.Sprint(v.Data[field])
	if !govalidator.IsEmail(value) {
		v.AddError(field, "邮箱地址无效")
	}
}

// Includes 是否包含某个值
func (v *JsonValidator) Includes(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}
