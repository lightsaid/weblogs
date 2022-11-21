{{template "desktop" .}}

{{define "title"}}首页{{end}}

{{define "style"}}
<link rel="stylesheet" href="/static/styles/index.page.css">
{{end}}

{{define "content"}}
<div class="gox-body-container">
    <!-- 主内容 -->
    <div class="gox-body-content">
        <div class="container-fluid">
            <h1 id="hello">Goxlog Hello!</h1>
            <form action="/user/update/1" method="post">
                <input type="hidden" name="csrf_token" value="{{ .CSRFToken }}">
                <button type="submit" class="btn login-btn btn-primary mt-1 mb-2">update</button>
            </form>
        </div>
    </div>

    <!-- Slider 右侧边栏 -->
    <div class="gox-body-slider">
        <div class="container-fluid sticky-top">
            <h1>Goxlog Slider!</h1>
            <h1>Goxlog Slider!</h1>
        </div>
    </div>
</div>
{{end}}

<!-- javascript -->
{{define "javascript"}}
<script src="/static/scripts/index.page.js"></script>
{{end}}