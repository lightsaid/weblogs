{{template "desktop" .}}


{{define "desktop-title"}}{{$post := index .DataMap "post"}}{{$post.Title}}{{end}}

{{define "desktop-style"}}
<link rel="stylesheet" href="/static/node_modules/editor.md/css/editormd.min.css" />
<link rel="stylesheet" href="/static/styles/detail.page.css">
{{end}}

{{define "desktop-content"}}
{{$post := index .DataMap "post"}}
<div class="gox-body-container">
    <!-- 主内容 -->
    <div class="gox-body-content">
        <div class="container-fluid">
            <div id="post-content"></div>
        </div>
    </div>

    <!-- Slider 右侧边栏 -->
    <div class="gox-body-slider">
        <div class="container-fluid sticky-top">
            <div id="sidebar">
                <h3>目录</h3>
                <div class="markdown-body editormd-preview-container table-content" id="toc"></div>
            </div>
        </div>
    </div>
</div>
{{end}}

<!-- javascript -->
{{define "desktop-javascript"}}
{{$post := index .DataMap "post"}}
<script type="text/javascript" src="/static/node_modules/editor.md/lib/marked.min.js"></script>
<script type="text/javascript" src="/static/node_modules/editor.md/lib/prettify.min.js"></script>
<script type="text/javascript" src="/static/node_modules/editor.md/lib/raphael.min.js"></script>
<script type="text/javascript" src="/static/node_modules/editor.md/lib/underscore.min.js"></script>
<script type="text/javascript" src="/static/node_modules/editor.md/lib/sequence-diagram.min.js"></script>
<script type="text/javascript" src="/static/node_modules/editor.md/lib/flowchart.min.js"></script>
<script type="text/javascript" src="/static/node_modules/editor.md/lib/jquery.flowchart.min.js"></script>
<script type="text/javascript" src="/static/node_modules/editor.md/editormd.min.js"></script>

<script type="text/javascript">
    $(function () {
        var contentEditor, catalogEditor;

        contentEditor = editormd.markdownToHTML("post-content", {
            markdown: "{{$post.Content}}",//markdown ,//+ "\r\n" + $("#append-test").text(),
            //htmlDecode      : true,       // 开启 HTML 标签解析，为了安全性，默认不开启
            htmlDecode: "style,script,iframe",  // you can filter tags decode
            //toc             : false,
            tocm: true,    // Using [TOCM]
            tocContainer: "#toc", // 自定义 ToC 容器层
            //gfm             : false,
            //tocDropdown     : true,
            // markdownSourceCode : true, // 是否保留 Markdown 源码，即是否删除保存源码的 Textarea 标签
            emoji: true,
            taskList: true,
            tex: true,  // 默认不解析
            flowChart: true,  // 默认不解析
            sequenceDiagram: true,  // 默认不解析
        });
    });
</script>
{{end}}