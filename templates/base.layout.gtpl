{{define "base"}}
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>轻言博客</title>
    <link rel="shortcut icon" href="/static/images/default/shortcut.png" type="image/x-icon">
    <link rel="stylesheet" href="/static/styles/root.css">
    {{block "style" .}} {{end}}
</head>
<body>
    <!-- <div id="toast" class="toast">Toast</div> -->

    {{block "content" .}} {{end}}

    {{block "javascript" .}} {{end}}

    <script src="/static/bootstrap/js/bootstrap.bundle.min.js"></script>
    <script>
        var toast = document.getElementById("toast")
        
        "{{with .Error}}"
            showToast("{{.}}","error")
        "{{end}}"

        "{{with .Success}}"
            showToast("{{.}}","success")
        "{{end}}"

        function showToast(text, className) {
            toast.innerText = text
            toast.classList.add(className)
            let timer = setTimeout(function(){
                toast.classList.remove(className)
                clearTimeout(timer)
            }, 2000)
        }
    </script>

</body>
</html>

{{end}}