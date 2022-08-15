package models

// TemplateData 定义模板数据结构
type TemplateData struct {
	StringMap map[string]string
	Data      map[string]interface{} // 数据
	RunMode   string                 // 环境
	Title     string                 // 页面标题
	Flash     string                 // info 消息
	Warning   string                 // 警告
	Error     string                 // 错误提示
}

// NewTemplateData 初始化一个TemplateData, 提供模板数据
func NewTemplateData() TemplateData {
	stringMap := make(map[string]string)
	data := make(map[string]interface{})
	return TemplateData{
		StringMap: stringMap,
		Data:      data,
		Title:     "",
		Warning:   "",
		Error:     "",
	}
}
