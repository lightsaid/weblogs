package validator

// 根据字段存放错误信息（提示）
type errors map[string][]string

// Add 根据字段添加错误西信息
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get 根据字段返回第一个错误信息
func (e errors) Get(field string) string {
	msgs := e[field]
	if len(msgs) == 0 {
		return ""
	}
	return msgs[0]
}
