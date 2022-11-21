{{template "base" .}}

{{define "title"}}登录{{end}}

{{define "style"}}
<link rel="stylesheet" href="/static/styles/login.page.css">
{{end}}


{{define "content"}}
{{$email := index .StringMap "email"}}
<main id="app">
    <div class="container-fluid">
        <div class="row row-cols-1 row-cols-md-2 g-3">
            <div class="col d-flex flex-column justify-content-center p-left-15">
                <h1 class="">轻言博客</h1>
                <h4>学以致用，坚持学习一百年不动摇</h4>
            </div>
            <div class="col">
                <form action="/user/login" method="POST" class="form-signin border rounded shadow" novalidate>
                    <h3 class="form-title fs-4 text-center fw-normal mb-3">欢迎登录</h3>
                    <input type="hidden" name="csrf_token" value="{{ .CSRFToken }}">
                    <div class="form-item mb-1">
                        <label for="email" class="form-label small">邮箱：</label>
                        <input type="email" name="email" class="form-control small  {{with .Form.Errors.Get "email"}} is-invalid {{end}}" value="{{$email}}" id="email" aria-describedby="emailHelp">
                        <span class="icon iconfont icon-user"></span>
                        <div id="emailHelp" class="form-text text-danger">
                            {{with .Form.Errors.Get "email"}} {{.}} {{end}}
                        </div>
                    </div>
                    <div class="form-item mb-1">
                        <label for="password" class="form-label small">密码：</label>
                        <input type="password" name="password" class="form-control small  {{with .Form.Errors.Get "password"}} is-invalid {{end}}" id="password" aria-describedby="pwsdHelp">
                        <span class="icon iconfont icon-password"></span>
                        <div id="pwsdHelp" class="form-text text-danger">
                            {{with .Form.Errors.Get "password"}} {{.}} {{end}}
                        </div>
                    </div>

                    <div class="form-item mb-1 d-flex justify-content-between align-items-center">
                        <div class="col-6 form-check">
                            <input type="checkbox" class="form-check-input" id="rememberme">
                            <label class="form-check-label small" for="rememberme">记住我</label>
                        </div>
                        <div class="col-6">
                            <a class="btn btn-link small"> 忘记密码？</a>
                        </div>
                    </div>
                    <div class="d-grid">
                        <button type="submit" class="btn login-btn btn-primary mt-1 mb-2">登录</button>
                    </div>
                </form>
            </div>

        </div>
    </div>
</main>

{{end}}