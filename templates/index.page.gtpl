{{template "desktop" .}}

{{define "desktop-title"}}首页{{end}}

{{define "desktop-style"}}
<link rel="stylesheet" href="/static/styles/index.page.css">
{{end}}

{{define "desktop-content"}}
{{$posts := index .DataMap "posts"}}
<div class="gox-body-container">
    <!-- 主内容 -->
    <div class="gox-body-content">
        <div class="container-fluid">
            <ul class="list-group list-group-flush">
                {{range $posts}}
                    <a href="/post/detail/{{.ID}}">
                        <li class="list-group-item">{{.Title}}</li>
                    </a>
                {{end}}
            </ul>
        </div>
    </div>

    <!-- Slider 右侧边栏 -->
    <div class="gox-body-slider">
        <div class="container-fluid sticky-top">

        </div>
    </div>
</div>
{{end}}

<!-- javascript -->
{{define "desktop-javascript"}}
<script src="/static/scripts/index.page.js"></script>
{{end}}