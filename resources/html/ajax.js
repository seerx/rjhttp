function ajaxGet() {
    return new Promise(function(resolve, reject) {
        var xhr = new XMLHttpRequest()
        xhr.open("GET", "/rj?m=graph", true)
        xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
        xhr.onreadystatechange = function () {
            if (xhr.readyState == 4 && xhr.status == 200) {
                resolve(xhr.responseText);
            }
        }
        xhr.send()
        // xhr.send("game=1&shang=common")
    })
}
