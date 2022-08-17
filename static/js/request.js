~function ($$) {
    var request = (url, options) => {
        let headers = new Headers()
        headers.append('Content-Type', 'application/json')
        var param = {
            method: options.method ? options.method : "POST",
            headers: Object.assign(headers, options.headers),
            body: JSON.stringify(options.body)
        }

        return fetch(url, param)
            .then(response => {
                let flag = false
                Object.values(response.headers).forEach(val => {
                    if (typeof val == "string") {
                        if (val.toLowerCase() == "application/json") {
                            flag = true
                        }
                    }
                })
                if (flag) {
                    return response.json()
                } else {
                    return response.text()
                }
            })
            .then(data => {
                if ($$.isProd !== "prod") {
                    // console.info(">> Request: ", url, "Response: ", data)
                }

                // 不会执行 <script> 标签里的代码 
                // let page = document.querySelector("html")
                // page.innerHTML = data

                // 执行script标签js, 需要将 const 定义的全局变量改成var/let, 
                // 另一种方式使用 document.createElement('script')
                document.write("")
                document.write(data)
                // 参考 https://developer.mozilla.org/zh-CN/docs/Web/API/Document/write
                document.close() // 必须close, 若close，第二次write就会空白。

                // 总结，这种混合开发模式有点意思~
            })
            .catch((err) => {
                console.error(err)
            })
    }
    $$.request = request
}(window)


