import axios from 'axios'

let service = axios.create({
    baseURL:  'http://127.0.0.1:517/api/v1/',
    timeout: 10000,
    headers: {
        'Content-type': 'application/josn'
    },
})
export default {
    //get请求，其他类型请求复制粘贴，修改method
    get(url, param, data) {
        return new Promise((cback, reject) => {
            service({
                method: 'get',
                url,
                params: param,
                data
            }).then(res => {
                //axios返回的是一个promise对象
                var res_code = res.status.toString();
                if (res_code.charAt(0) == 2) {
                    cback(res); //cback在promise执行器内部
                } else {
                    console.log(res, '异常1')
                }
            }).catch(err => {
                if (!err.response) {
                    console.log('请求错误')
                } else {
                    reject(err.response);
                    console.log(err.response, '异常2')
                }
            })

        })
    },
    delete(url, param) {
        return new Promise((cback, reject) => {
            service({
                method: 'delete',
                url,
                params: param,
            }).then(res => {
                //axios返回的是一个promise对象
                var res_code = res.status.toString();
                if (res_code.charAt(0) == 2) {
                    cback(res); //cback在promise执行器内部
                } else {
                    console.log(res, '异常1')
                }
            }).catch(err => {
                if (!err.response) {
                    console.log('请求错误')
                } else {
                    reject(err.response);
                    console.log(err.response, '异常2')
                }
            })

        })
    },
    post(url, param, data) {
        return new Promise((cback, reject) => {
            service({
                method: 'post',
                url,
                params: param,
                data,
            }).then(res => {
                //axios返回的是一个promise对象
                var res_code = res.status.toString();
                if (res_code.charAt(0) == 2) {
                    cback(res); //cback在promise执行器内部
                } else {
                    console.log(res, '异常1')
                }
            }).catch(err => {
                if (!err.response) {
                    console.log('请求错误')
                } else {
                    reject(err.response);
                    console.log(err.response, '异常2')
                }
            })

        })
    },
    put(url, param, data) {
        return new Promise((cback, reject) => {
            service({
                method: 'put',
                url,
                params: param,
                data,
            }).then(res => {
                //axios返回的是一个promise对象
                var res_code = res.status.toString();
                if (res_code.charAt(0) == 2) {
                    cback(res); //cback在promise执行器内部
                } else {
                    console.log(res, '异常1')
                }
            }).catch(err => {
                if (!err.response) {
                    console.log('请求错误')
                } else {
                    reject(err.response);
                    console.log(err.response, '异常2')
                }
            })

        })
    }
}
