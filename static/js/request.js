~function ($$) {
    var request = (url, options) => {
        let headers = new Headers()
        headers.append('Content-Type', 'application/json')
        const param = {
            method: options.method ? options.method : "POST",
            headers: Object.assign(headers, options.headers),
            body: JSON.stringify(options.body)
        }

        return fetch(url, param)
            .then(response => response.json())
            .then(data => {
                if ($$.isProd !== "prod") {
                    console.info(">> Request: ", url, "Response: ", data)
                }
            })
            .catch((error) => {
                console.error(error)
            })
    }
    $$.request = request
}(window)

