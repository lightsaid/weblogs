package forms

type errors map[string][]string

// Add 根据字段添加错误信息
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get 从 errors 根据 field 取出第一个错误，如果有的话
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
