{{define "header"}}
    <!-- Header 顶部导航栏组件 -->
    <div class="container-fluid gox-header">
        <nav class="navbar fixed-top navbar-expand navbar-light bg-light">
            <div class="container gox-header-navbar">
                <div class="collapse navbar-collapse" id="navbar-collapse">
                    <a class="navbar-brand" href="#">
                        <!-- <img src="/static/images/default/goxlog01.svg" class="logo" alt="logo"> -->
                    </a>
                   
                    <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                        <li class="nav-item home  {{if eq .RequestPath "/"}}active{{end}}">
                            <a class="nav-link" aria-current="page" href="/"> <i class="iconfont icon-home-wifi-fill"></i>Home</a>
                        </li>
                        <li class="nav-item notifications {{if eq .RequestPath "/notification"}}active{{end}}">
                            <a class="nav-link" href="/notification"> <i class="iconfont icon-a-5Hzhenling"></i>Notifications</a>
                        </li>
                        <li class="nav-item tags {{if eq .RequestPath "/tag"}}active{{end}}">
                            <a class="nav-link" href="/tag"> <i class="iconfont icon-tags"></i>Tags</a>
                        </li>
                        <li class="nav-item about {{if eq .RequestPath "/user/about"}}active{{end}}">
                            <a class="nav-link" href="/user/about"> <i class="iconfont icon-guanyuwo"></i>About Me</a>
                        </li>
                    </ul>

                    <form class="d-flex">
                        <input class="form-control me-2 small" type="search" placeholder="Search" aria-label="Search">
                    </form>

                    <li class="online">
                        <div class="dropdown">
                          <button class="btn btn-sm btn-white dropdown-toggle user-meun" type="button" id="dropdownMenuButton"
                            data-bs-toggle="dropdown" aria-expanded="false">
                            <img class="img-fluid avatar" src="/static/images/default/avatar3.png" alt="">
                            Lightsaid
                          </button>
                          <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                            <li>
                              <a class="dropdown-item new-post" href="">
                                <i class="iconfont icon-write"></i>
                                <span>New Posts</span>
                              </a>
                            </li>
                            <li>
                              <a class="dropdown-item" href="/user/settings">
                                <i class="iconfont icon-shezhi"></i>
                                <span>Settings</span>
                              </a>
                            </li>
                            <li>
                              <a class="dropdown-item" href="/user/logout">
                                <i class="iconfont icon-tuichu"></i>
                                <span>Log Out</span>
                              </a>
                            </li>
                          </ul>
                        </div>
                      </li>

                    <!-- 外部链接 -->
                    <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                        <li class="nav-item github">
                            <a class="nav-link" href="#"> <i class="iconfont icon-github"></i>Github</a>
                        </li>
                        <li class="nav-item bilibili">
                            <a class="nav-link" href="#"><i class="iconfont icon-bilibili"></i></a>
                        </li>
                        <li class="nav-item mode">
                          <!-- <a class="nav-link" href="#"><i class="iconfont icon-Daytimemode"></i>Day</a> -->
                          <!-- <a class="nav-link" href="#"><i class="iconfont icon-nightmode-fill"></i>Night</a> -->
                      </li>
                    </ul>
                </div>
            </div>
        </nav>
    </div>

{{end}}