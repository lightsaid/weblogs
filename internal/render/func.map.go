package render

import (
	"fmt"
	"html/template"
	"os"

	"lightsaid.com/weblogs/internal/service"
)

func imageURL(url string) string {
	if len(url) > 2 && url[0] == '.' {
		prefix := os.Getenv("ASSETS_PREFIX")
		return fmt.Sprintf("%s%s", prefix, url[1:])
	}
	return url
}

func ifAdminF(status int) string {
	if status == 1 {
		return "是"
	}
	return "否"
}

// 状态 (-1:删除0:正常|1:活跃)
func getActive(active int) string {
	switch active {
	case -1:
		return "已删除"
	case 0:
		return "正常"
	case 1:
		return "活跃"
	default:
		return fmt.Sprintf("未知%d", active)
	}
}

func getAttrKind(kind string) string {
	if kind == "T" {
		return "标签(Tag)"
	}
	return "标记(Mark)"
}

// 从二级开始递归元素
func recursionCategory(categories []*service.LevelCategories) string {
	// NOTE: 从二级开始，ul 元素有 children class
	var html = `<ul class="list-group children">`

	for _, v := range categories {
		html += fmt.Sprintf(`<li class="list-group-item">
		<input class="form-check-input me-1" type="checkbox" value="%d" id="%d">
		<label class="form-check-label stretched-link" for="%d">%s</label>`,
			v.Category.ID, v.Category.ID, v.Category.ID, v.Category.Name)
		if len(v.Children) > 0 {
			// 递归2、3、4...子分类 DOM
			html += recursionCategory(v.Children)
		} else {
			// 没有子元素了，闭合标签
			html += "</li>"
		}
	}

	html += "</ul>"
	return html
}

// 递归渲染分类
func renderCheckboxCategories(categories []*service.LevelCategories) template.HTML {
	// 顶级 ul 元素
	var html = `<ul class="list-group">`

	// 遍历一级分类
	for _, v := range categories {
		html += fmt.Sprintf(`<li class="list-group-item">
				<input class="form-check-input me-1" type="checkbox" value="%d" id="%d">
				<label class="form-check-label stretched-link" for="%d">%s</label>`,
			v.Category.ID, v.Category.ID, v.Category.ID, v.Category.Name)
		if len(v.Children) > 0 {
			html += recursionCategory(v.Children)
		} else {
			html += "</li>"
		}
	}

	html += "</ul>"

	// NOTE: 分类显示
	// 直接返回 html string, html/template 为了安全是不会渲染的，按普通文本输出显示
	// return html

	// Nice 显示正常 （参考: https://go.dev/play/p/U64_7UHZQU）
	return template.HTML(html)

	/*
		{{range $v := $categories}}
			<li class="list-group-item">
				<input class="form-check-input me-1" type="checkbox" value="{{$v.Category.ID}}" id="{{$v.Category.ID}}">
				<label class="form-check-label stretched-link" for="{{$v.Category.ID}}">{{$v.Category.Name}}</label>

				<!-- 二级分类 -->
				{{$length := len $v.Children}}
				{{if gt $length 0}}

					<ul class="list-group children">
						{{range $vv := $v.Children}}
						<li class="list-group-item">
							<input class="form-check-input me-1" type="checkbox" value="{{$vv.Category.ID}}" id="{{$vv.Category.ID}}">
							<label class="form-check-label stretched-link" for="{{$vv.Category.ID}}">{{$vv.Category.Name}}</label>
						</li>
						{{end}}
					</ul>f
				{{end}}
			</li>
		{{end}}
	*/
}
