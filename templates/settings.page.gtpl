{{template "desktop" .}}

{{define "title"}}Settings{{end}}

{{define "style"}}
<link rel="stylesheet" href="/static/styles/settings.page.css">
{{end}}

{{define "content"}}

<div class="gox-body-content">
    <div class="container-fluid">
        <h4 class="title position-relative">Settings</h4>
        <!-- Tabs -->
        <nav>
            <div class="nav nav-tabs mb-3" id="nav-tab" role="tablist">
                <button class="nav-link active" id="nav-profile-tab" data-bs-toggle="tab" data-bs-target="#nav-profile"
                    type="button" role="tab" aria-controls="nav-profile" aria-selected="true">Profile</button>
                <button class="nav-link" id="nav-password-tab" data-bs-toggle="tab" data-bs-target="#nav-password"
                    type="button" role="tab" aria-controls="nav-password" aria-selected="false">Password</button>
            </div>
        </nav>
        <div class="tab-content" id="nav-tabContent">
            <!-- Profile -->
            <div class="tab-pane fade show active" id="nav-profile" role="tabpanel" aria-labelledby="nav-profile-tab">
                <form class="form">
                    <div class="mb-3">
                        <label for="email" class="form-label">Email address</label>
                        <input type="email" class="form-control" id="email" value="Lightsaid" disabled aria-describedby="emailHelp">
                    </div>
                    <div class="mb-3">
                        <label for="username" class="form-label">UserName</label>
                        <input type="username" class="form-control" id="username" value="Lightsaid" aria-describedby="usernameHelp">
                        <p id="usernameHelp" class="form-text small">用户名长度在2和16之间</p>
                    </div>
                    <div class="mb-3">
                        <label class="form-label" for="avatarFile">Avatar</label>
                        <div class="profile-avatar">
                            <div class="profile-avatar-container">
                                <img src="/static/images/default/avatar3.png" class="img-thumbnail rounded-circle">
                                <div class="upload-icon"> <i class="iconfont icon-tianjiatupian_huaban"></i> </div>
                            </div>
                        </div>
                        <input type="file" hidden class="form-control" name="avatarFile" id="avatarFile">
                        <p class="form-text small">头像图片格式必须是 ".png / .jpg / .jpeg" </p>
                    </div>
                    <button type="submit" class="btn theme-btn save-btn btn-primary">Save</button>
                </form>
            </div>

            <!-- Password -->
            <div class="tab-pane fade" id="nav-password" role="tabpanel" aria-labelledby="nav-password-tab">
                Password Manager
            </div>
        </div>
    </div>
</div>
</div>

{{end}}

{{define "javascript"}}
<script>
    $(function(){
        $(".profile-avatar-container").click(function(){
            $("input[name=avatarFile]").click()
        })
    })
</script>
{{end}}