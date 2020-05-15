
function Ajax(rootUrl, headers) {
    let self = this
    self.root = rootUrl
    self.headers = headers

    this.Upload = function (param, file, fieldName, tokenInHeader, tokenInCookie, tokenInSetCookie) {
        return new Promise(function(resolve, reject) {
            let xhr = new XMLHttpRequest()
            xhr.open("POST", self.root, true)
            // xhr.setRequestHeader('Content-type', 'multipart/form-data');
            if (self.headers) {
                for (let k in self.headers) {
                    if (tokenInSetCookie) {
                        xhr.setRequestHeader("Set-Cookie", k + "=" + self.headers[k])
                    }
                    if (tokenInCookie) {
                        xhr.setRequestHeader("Cookie", k + "=" + self.headers[k])
                    } 
                    if (tokenInHeader) {
                        xhr.setRequestHeader(k, self.headers[k]);
                    }
                }
            }
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

    this.GetX = function(param, tokenInHeader, tokenInCookie, tokenInSetCookie) {
        let openImage = function (res) {
            let win = window.open('about:blank')
            with (win.document) {
                var img = createElement('img')
                body.setAttribute('style', 'margin: 0;')
                img.setAttribute('style', 'width: auto; height: 100%;')
                img.src = window.URL.createObjectURL(res)
                img.onload = function(){
                    //图片加载完，释放一个URL资源。
                    window.URL.revokeObjectURL(this.src)
                }
                body.appendChild(img)
            }
        }
        let saveFile = function (xhr) {
            // console.log(xhr.getAllResponseHeaders())
            let name = xhr.getResponseHeader("Content-disposition");
            let filename = name
            if (filename && filename.length > 0) {
                let fb = filename.indexOf('=')
                if (fb >= 0) {
                    filename = filename.substring(fb + 1, filename.length);
                }
            }
            // console.log(xhr.response)
            let blob = new Blob([xhr.response])
            let url = URL.createObjectURL(blob);
            let link = document.createElement('a');
            link.href = url
            if (filename && filename.length > 0) {
                link.download = filename
            }
            link.click()
        }
        return new Promise(function(resolve, reject) {
            let xhr = new XMLHttpRequest()
            let url = self.root + '?' + param
            xhr.open("GET", url, true)
            if (self.headers) {
                for (let k in self.headers) {
                    if (tokenInSetCookie) {
                        xhr.setRequestHeader("Set-Cookie", k + "=" + self.headers[k])
                    }
                    if (tokenInCookie) {
                        xhr.setRequestHeader("Cookie", k + "=" + self.headers[k])
                    } 
                    if (tokenInHeader) {
                        xhr.setRequestHeader(k, self.headers[k]);
                    }
                }
            }
            xhr.responseType = 'blob'
            // xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
            xhr.onreadystatechange = function () {
                if (xhr.readyState == 4 && xhr.status == 200) {
                    resolve(xhr.response);
                    // console.log('xhr.responseType', xhr)
                    if (xhr.response.type.indexOf('image/') === 0) {
                        openImage(xhr.response)
                    } else {
                        saveFile(xhr)
                    }
                }
            }
            xhr.onerror =function (e) {
                resolve('** An error occurred during the transaction')
            };
            try {
                xhr.send()
            } catch (e) {
                resolve(e)
            }
            // xhr.send("game=1&shang=common")
        })
    }

    this.Get = function(param, tokenInHeader, tokenInCookie, tokenInSetCookie) {
        return new Promise(function(resolve, reject) {
            let xhr = new XMLHttpRequest()
            let url = self.root + '?' + param
            xhr.open("GET", url, true)
            if (self.headers) {
                for (let k in self.headers) {
                    if (tokenInSetCookie) {
                        xhr.setRequestHeader("Set-Cookie", k + "=" + self.headers[k])
                    }
                    if (tokenInCookie) {
                        xhr.setRequestHeader("Cookie", k + "=" + self.headers[k])
                    } 
                    if (tokenInHeader) {
                        xhr.setRequestHeader(k, self.headers[k]);
                    }
                }
            }
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

    this.Post = function(param, tokenInHeader, tokenInCookie, tokenInSetCookie) {
        return new Promise(function(resolve, reject) {
            let xhr = new XMLHttpRequest()
            // let url = self.root + '?' + param
            xhr.open("POST", self.root, true)
            if (self.headers) {
                for (let k in self.headers) {
                    if (tokenInSetCookie) {
                        xhr.setRequestHeader("Set-Cookie", k + "=" + self.headers[k])
                    }
                    if (tokenInCookie) {
                        xhr.setRequestHeader("Cookie", k + "=" + self.headers[k])
                    } 
                    if (tokenInHeader) {
                        xhr.setRequestHeader(k, self.headers[k]);
                    }
                }
            }
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
