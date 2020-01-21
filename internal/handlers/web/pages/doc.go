package pages

/**
发布包时，运行 scripts/gen_static.sh 可以生成 static.go 文件

调试 html 页面时，运行 scripts/gen_static_debug.sh 脚本

或者使用下面的方式
生成 static.go 文件
在项目根目录，执行此命令
go-bindata -pkg=pages -ignore=./.idea -o ../../pkg/handlers/web/pages/static.go ...

go-binddata 命令
go get -u github.com/jteeuwen/go-bindata/...
*/
