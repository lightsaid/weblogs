{{template "base" .}}

{{define "title"}}New Post{{end}}

{{define "style"}}
<link rel="stylesheet" href="/static/node_modules/editor.md/css/editormd.min.css" />
<link rel="stylesheet" href="/static/styles/create.post.page.css">
{{end}}

{{define "content"}}

<div class="container-fluid fixed-top">
    <ul class="d-flex justify-between header">
        <li class="nav-item title">
            <input type="text" class="form-control" id="title" placeholder="Enter the post title...">
        </li>
        <li class="nav-item submit-btns">
            <button class="btn btn-secondary dropdown-toggle" id="choose-tag" 
                type="button" id="dropdownMenuButton" data-bs-toggle="dropdown">
                Choose Tags
            </button>

            <div class="btn-group" role="group" id="publish" aria-label="Basic">
                <button class="btn small theme-btn btn-primary">
                    <i class="iconfont icon-fasong"></i>
                    Publish
                </button>
                <button class="btn small save-draft" id="save-draft"><i class="iconfont icon-ziliao"></i>Save Draft</button>
            </div>
        </li>
    </ul>
</div>

<div class="offcanvas offcanvas-end" tabindex="-1" id="offcanvas" aria-labelledby="offcanvasExampleLabel">
    <div class="offcanvas-header">
        <h5 class="offcanvas-title" id="offcanvasExampleLabel">Choose Tags</h5>
        <button type="button" class="btn-close" id="offcanvas-close" data-bs-dismiss="offcanvas" aria-label="Close"></button>
    </div>
    <div class="offcanvas-body">
        <form action="#" class="tag-form">
            <ul class="list-group">
                <li class="list-group-item">
                <input class="form-check-input me-1" name="tag" type="checkbox" value="1">
                First checkbox
                </li>
                <li class="list-group-item">
                <input class="form-check-input me-1" name="tag" type="checkbox" value="2">
                Second checkbox
                </li>
          </ul>
        </form>
    </div>
</div>
<div class="offcanvas-backdrop fade show" style="display: none;"></div>

<div style="margin-top:66px">
    <div id="post-editor"></div>
</div>
{{end}}

{{define "javascript"}}
<script src="/static/scripts/jquery.min.js"></script>
<script src="/static/node_modules/editor.md/editormd.min.js"></script>

<script type="text/javascript">
    // editormd.md 和 bootstrap.bundle.js 冲突了，无效对offcanvas组件工具，因此手动实现
    let chooseTag = document.querySelector("#choose-tag")
    let offcanvas = document.querySelector("#offcanvas")
    let offcanvasClose = document.querySelector("#offcanvas-close")
    let offcanvasBackdrop = document.querySelector(".offcanvas-backdrop")
    let publish = document.querySelector("#publish")
    let saveDraft = document.querySelector("#save-draft")

    chooseTag.addEventListener("click", function(){
        offcanvasBackdrop.style.display = "block"
        offcanvas.classList.toggle("show")
        offcanvas.style.transform = "none"
    })

    offcanvasClose.addEventListener("click", function(){
        offcanvasBackdrop.style.display = "none"
        offcanvas.style.transform = "translateX(100%)"
        let timer =  setTimeout(function(){
            offcanvas.classList.toggle("show")
            clearTimeout(timer)
        }, 500)
    })

    publish.addEventListener("click", function(){
        let data = getData()
        data.active = 1
        sendData(data)
    })

    saveDraft.addEventListener("click", function(){
        let data = getData()
        data.active = 0
        sendData(data)
    })

    function sendData(data){
        $.ajax({url:"/post/create",type:"POST",data: JSON.stringify(data), dataType:"json",success: function(response){
            // TODO:
            console.log(response)
            if (!response.ok){
                alert(response.error)
                return
            }
            location.href = "/"
        }})
    }

    function getData() {
        let title = getTitle()
        let tagIDs = getChooseTags()
        let postContent = posteditor.getMarkdown()
        return {
            title: title,
            tag_ids: tagIDs,
            content: postContent,
            token: "{{.CSRFToken}}",
            active: 1
        }
    }

    function getTitle(){
        return document.querySelector("#title").value
    }

    function getChooseTags() {
        let form = document.querySelector(".tag-form")
        let tagIDs = []
        form.elements.tag.forEach(function(tag){
            if (tag.checked) {
                tagIDs.push(parseInt(tag.value))
            }
        })
        return tagIDs
    }

    var posteditor;
    $(function () {
        var testMDUrl = "/static/node_modules/editor.md/examples/test.md"
        $.get(testMDUrl, function (md) {
            posteditor = editormd("post-editor", {
                width: "99%",
                height: "calc(100vh - 80px)",
                path: '/static/node_modules/editor.md/lib/',
                theme: "white",
                previewTheme: "white",
                editorTheme: "default",
                markdown: md,
                codeFold: true,
                //syncScrolling : false,
                saveHTMLToTextarea: true,    // 保存 HTML 到 Textarea
                searchReplace: true,
                //watch : false,                // 关闭实时预览
                htmlDecode: "style,script,iframe|on*",            // 开启 HTML 标签解析，为了安全性，默认不开启    
                //toolbar  : false,             //关闭工具栏
                //previewCodeHighlight : false, // 关闭预览 HTML 的代码块高亮，默认开启
                emoji: true,
                taskList: true,
                tocm: true,         // Using [TOCM]
                tex: true,                   // 开启科学公式TeX语言支持，默认关闭
                flowChart: true,             // 开启流程图支持，默认关闭
                sequenceDiagram: true,       // 开启时序/序列图支持，默认关闭,
                //dialogLockScreen : false,   // 设置弹出层对话框不锁屏，全局通用，默认为true
                //dialogShowMask : false,     // 设置弹出层对话框显示透明遮罩层，全局通用，默认为true
                //dialogDraggable : false,    // 设置弹出层对话框不可拖动，全局通用，默认为true
                //dialogMaskOpacity : 0.4,    // 设置透明遮罩层的透明度，全局通用，默认值为0.1
                //dialogMaskBgColor : "#000", // 设置透明遮罩层的背景颜色，全局通用，默认为#fff
                imageUpload: true,
                imageFormats: ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
                imageUploadURL: "/upload",
                onload: function () {
                    console.log('onload', this);
                    //this.fullscreen();
                    //this.unwatch();
                    //this.watch().fullscreen();

                    //this.setMarkdown("#PHP");
                    //this.width("100%");
                    //this.height(480);
                    //this.resize("100%", 640);
                }
            });
        });
    });
</script>

{{end}}