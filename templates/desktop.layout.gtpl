{{define "desktop"}}
<!DOCTYPE html>
<html lang="zh-CN">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{template "title" .}}</title>
    <link rel="shortcut icon" href="/static/images/default/shortcut.png" type="image/x-icon">
    <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/styles/font.css">
    <link rel="stylesheet" href="/static/fonts/iconfont/iconfont.css">
    <link rel="stylesheet" href="/static/styles/root.css">
    <link rel="stylesheet" href="/static/styles/header.partial.css">

    {{block "style" .}} {{end}}

    <style>
        .gox-body {
            width: 1024px;
            max-width: 1024px;
            min-height: calc(100vh - 56px);
            margin: 0 auto;
            position: relative;
            padding-top: 56px;
        }

        .gox-body::before {
            display: block;
            content: " ";
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            z-index: -1;
            background: var(--canvas);
        }
    </style>
</head>

<body>
    <!-- Header 顶部导航栏 -->
    {{template "header" .}}

    <div class="container-fluid gox-body">
        {{block "content" .}} {{end}}
    </div>

    <script src="/static/bootstrap/js/bootstrap.bundle.min.js"></script>
    <script src="/static/scripts/jquery.min.js"></script>

    {{block "javascript" .}} {{end}}

</body>

</html>

{{end}}