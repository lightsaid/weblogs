{{template "desktop" .}}

{{define "desktop-title"}}Tags{{end}}

{{define "desktop-style"}}
<link rel="stylesheet" href="/static/styles/tags.page.css">
{{end}}

{{define "desktop-content"}}
{{$tags := index .DataMap "tags"}}
<div class="gox-body-content">
    <div class="container-fluid">
        <div class="header position-relative"> 
            <h4 class="title position-relative">
                All Tags <span class="badge rounded-pill">{{len $tags}}</span>
            </h4>
            {{if .IsAuthenticated}}
            <button type="button" class="btn btn-primary btn-sm position-absolute new-tag" data-bs-toggle="modal"
                data-bs-target="#newTagModal"> New Tag </button>
            {{end}}
        </div>

        <ul class="list-group list-group-flush">
            {{range $tags}}
            <li class="list-group-item d-flex justify-content-between align-items-center">
                <a href="#">#{{.TagID}} - {{.Name}}</a>
                <span class="badge bg-primary rounded-pill">{{.TagCount}}</span>
            </li>
            {{end}}
        </ul>
    </div>
</div>

<!-- New Tag 弹窗 -->
<div class="modal fade" id="newTagModal" tabindex="-1" aria-labelledby="newTagModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="newTagModalLabel">New Tag</h5>
              </div>
            <div class="modal-body">
                <form action="/tag/create" method="post">
                    <input type="hidden" name="csrf_token" value="{{ .CSRFToken }}">
                    <div class="mb-3">
                        <label for="tag" class="col-form-label">Tag Name:</label>
                        <input type="text" class="form-control" name="tag" id="tag">
                    </div>
                    <input type="submit" id="form-submit" value="submit" hidden>
                </form>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn small btn-danger" data-bs-dismiss="modal">Close</button>
                <button type="button" class="btn small btn-primary" id="tag-save">Save</button>
            </div>
        </div>
    </div>
</div>

{{end}}


{{define "desktop-javascript"}}
<script>
    var newTagModal = new bootstrap.Modal(document.getElementById('newTagModal'), {
        keyboard: false
    })
    
    $("#tag-save").on("click", function(){
        $("#form-submit").click();
        newTagModal.toggle();
    })
</script>
{{end}}