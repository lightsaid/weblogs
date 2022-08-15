~function ($$) {
    // 处理alert
    const alertPlaceholder = document.querySelector('.alert-container')
    const alert = (message, type) => {
        if (type == 'error') {
            type = 'danger'
        }
        const wrapper = document.createElement('div')
        wrapper.innerHTML = [
            `<div class="alert alert-${type} alert-dismissible" role="alert">`,
            `   <div class="text-center">${message}</div>`,
            '   <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>',
            '</div>',
        ].join('')

        alertPlaceholder.append(wrapper)

        let btns = alertPlaceholder.querySelectorAll(".btn-close")
        Array.from(btns).forEach((btn, index) => {
            setTimeout(() => {
                btn.click()
            }, 3000 + index * 1200);
        })

    }
    $$.showMsg = alert
}(window)
