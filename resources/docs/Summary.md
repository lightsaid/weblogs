# 总结

**写写总结，记录一下，走得更远，学得更快。**

<!-- TODO: 写总结、笔记 -->

### 注册
在 login.page.tmpl 页面尝试使用多种风格组合在一起，使用 Vue3 + Fetch + golang template
- Fetch 请求是一个异步请求，本身不像 Form 表单提交会刷新页面。
- 因此而是需要编码渲染，手动渲染要解决script标签代码不执行问题。
    - 整个页面的替换使用 innerHTML 是不会执行script里的代码的。
    - 因此更为简单的方式是：write前会默认执行open,wirte 完后，必须手动close，如若不close，再次write时会产生白屏。
        ``` js
            document.write("")
            document.write(domString)
            document.close()
        ```

- 当然fetch也能发送formData数据，会不会刷新页面呢？猜测：不会，毕竟是模拟。
- 另个既然使用异步请求，优先还是请求返回JOSN数据，然后前端直接操作数据较为方便。
- `注意`: 这是一个不合理的尝试，因为前端发送异步请求，后续逻辑跳转逻辑应由前端处理，而不是后台去重定向，不然后台虽重定向了，但是浏览器地址不会发生改变。(前端发起的异步请求，后台重定向不会改变浏览器地址)

- 使用引进Vue，未必是好事，也未必是坏事，视页面复杂程度。
- 由于 Go html/template 包使用 "{{}}" 占位符渲染模板和 Vue 有冲突，因此Vue不能再用， 可以使用指令v-html解决。

### Reflect


### gorilla/csrf


### gorilla/session
- session.AddFlash("message body", "Flash")
