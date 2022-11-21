{{template "desktop" .}}

{{define "title"}}Tags{{end}}

{{define "style"}}
<link rel="stylesheet" href="/static/styles/tags.page.css">
{{end}}

{{define "content"}}

<div class="gox-body-content">
    <div class="container-fluid">
        <div class="header position-relative">
            <h4 class="title position-relative">
                All Tags <span class="badge rounded-pill">99</span>
            </h4>
            <button type="button" class="btn btn-primary btn-sm position-absolute new-tag"> New Tag </button>
        </div>

        <ul class="list-group list-group-flush">
            <li class="list-group-item d-flex justify-content-between align-items-center">
                <a href="#">A list item</a> 
                <span class="badge bg-primary rounded-pill">14</span>
            </li>
            <li class="list-group-item d-flex justify-content-between align-items-center">
                <a href="#">A list item</a> 
                <span class="badge bg-primary rounded-pill">2</span>
            </li>
            <li class="list-group-item d-flex justify-content-between align-items-center">
                <a href="#">A list item</a> 
                <span class="badge bg-primary rounded-pill">1</span>
            </li>
        </ul>
    </div>
</div>

{{end}}