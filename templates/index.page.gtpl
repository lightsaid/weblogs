<!doctype html>
<html lang="zh-CN">

<head>
    <!-- 必须的 meta 标签 -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- Favicon 网站图标 TODO: 设计一个好看的，先随便放一个-->
    <link rel="shortcut icon" href="/static/images/default/shortcut.png" type="image/x-icon">

    <!-- Bootstrap 的 CSS 文件 -->
    <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css">

    <!-- 公共 CSS -->
    <link rel="stylesheet" href="/static/styles/font.css">

    <!-- 布局 CSS -->
    <link rel="stylesheet" href="/static/styles/index.page.css">

    <title>Goxlog</title>
</head>

<body>
    <!-- Header 顶部导航栏 -->
    <div class="container-fluid gox-header">
        <nav class="navbar fixed-top navbar-expand navbar-light bg-light">
            <div class="container gox-header-navbar">
                <div class="collapse navbar-collapse" id="navbar-collapse">
                    <a class="navbar-brand" href="#">
                        <img src="/static/images/default/goxlog01.svg" class="logo" alt="logo">
                    </a>
                    <form class="d-flex">
                        <input class="form-control me-2" type="search" placeholder="Search" aria-label="Search">
                    </form>
                    <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                        <li class="nav-item">
                            <a class="nav-link active" aria-current="page" href="#">Home</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="#">Notifications</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="#">Tags</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="#">About Me</a>
                        </li>
                    </ul>
                   
                </div>
            </div>
        </nav>
    </div>

    <!-- Body 主体 -->
    <div class="container-fluid gox-body">
        <div class="gox-body-container">
            <!-- Gox 微博内容 -->
            <div class="gox-body-content">
                <div class="container-fluid">
                    <h1>Goxlog Hello!</h1>
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
    </div>

    <!-- JavaScript：包含 Popper 的 Bootstrap 集成包 -->
    <script src="/static/bootstrap/js/bootstrap.bundle.min.js"></script>

    <!-- 逻辑处理的 JS -->
    <script src="/static/scripts/index.page.js"></script>
</body>

</html>