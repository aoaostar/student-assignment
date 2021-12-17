const instance = axios.create({
    baseURL: '/api/',
    // timeout: 1000,
    // headers: {'X-Custom-Header': 'foobar'}
});

const httpGet = (url) => {
    return instance.get(url)
        .then(res => {
            return res.data
        }).then(res => {
            if (res.status !== "ok") {
                mdui.snackbar({
                    message: res.message,
                    position: 'top',
                });
            }
            return res
        })
}
const httpPost = (url, params) => {
    return instance.post(url, params)
        .then(res => {
            return res.data
        }).then(res => {
            mdui.snackbar({
                message: res.message,
                position: 'top',
            });
            return res
        })
}
const httpPut = (url, params) => {
    return instance.put(url, params)
        .then(res => {
            return res.data
        }).then(res => {
            mdui.snackbar({
                message: res.message,
                position: 'top',
            });
            return res
        })
}
const httpDelete = (url, params) => {
    return instance.delete(url, {
        params
    })
        .then(res => {
            return res.data
        }).then(res => {
            mdui.snackbar({
                message: res.message,
                position: 'top',
            });
            return res
        })
}