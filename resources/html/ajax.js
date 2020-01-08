
function Ajax(rootUrl) {
    let self = this
    self.root = rootUrl

    this.Upload = function (param, file, fieldName) {
        return new Promise(function(resolve, reject) {
            let xhr = new XMLHttpRequest()
            xhr.open("POST", self.root, true)
            // xhr.setRequestHeader('Content-type', 'multipart/form-data');
            xhr.setRequestHeader('--run-json-field--', 'body');
            xhr.onreadystatechange = function () {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    resolve(xhr.responseText);
                }
            }
            xhr.onerror =function (e) {
                reject('** An error occurred during the transaction')
            };
            try {
                let form = new FormData()
                form.append('body', param)
                form.append(fieldName, file)
                xhr.send(form)
            } catch (e) {
                reject(e)
            }
        })
    }

    this.Get = function(param) {
        return new Promise(function(resolve, reject) {
            let xhr = new XMLHttpRequest()
            let url = self.root + '?' + param
            xhr.open("GET", url, true)
            xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
            xhr.onreadystatechange = function () {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    resolve(xhr.responseText);
                }
            }
            xhr.onerror =function (e) {
                reject('** An error occurred during the transaction')
            };
            try {
                xhr.send()
            } catch (e) {
                reject(e)
            }
            // xhr.send("game=1&shang=common")
        })
    }

    this.Post = function(param) {
        return new Promise(function(resolve, reject) {
            let xhr = new XMLHttpRequest()
            // let url = self.root + '?' + param
            xhr.open("POST", self.root, true)
            xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
            xhr.onreadystatechange = function () {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    resolve(xhr.responseText);
                }
            }
            xhr.onerror =function (e) {
                reject('** An error occurred during the transaction')
            };
            try {
                xhr.send(param)
            } catch (e) {
                reject(e)
            }
        })
    }
}

// function ajaxGet() {
//     return new Promise(function(resolve, reject) {
//         var xhr = new XMLHttpRequest()
//         xhr.open("GET", "/rj?m=graph", true)
//         xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
//         xhr.onreadystatechange = function () {
//             if (xhr.readyState == 4 && xhr.status == 200) {
//                 resolve(xhr.responseText);
//             }
//         }
//         xhr.onerror =function (e) {
//             reject(e)
//         };
//         xhr.send()
//         // xhr.send("game=1&shang=common")
//     })
// }
