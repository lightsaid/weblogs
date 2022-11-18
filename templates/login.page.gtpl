{{template "base" .}}

{{define "style"}}
<link rel="stylesheet" href="/static/styles/login.page.css">
{{end}}

{{define "content"}}
    <form action="/login" method="post" class="login-form">
        <input type="hidden" name="csrf_token" value="{{ .CSRFToken }}">
        <div class="form-item">
            <label for="email">Email:</label>
            <input type="text" name="email" autocomplete="off" id="email">
            {{with .Form.Errors.Get "email"}}
                <p class="error">{{.}}</p>
            {{end}}
        </div>

        <div class="form-item">
            <label for="password">Password:</label>
            <input type="password" name="password" id="password">
            {{with .Form.Errors.Get "password"}}
                <p class="error">{{.}}</p>
            {{end}}
        </div>

        <div class="form-item">
            <button type="submit" class="btn">登录</button>
        </div>
    </form>
{{end}}

