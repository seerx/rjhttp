#!/usr/bin/env bash
go get -u github.com/jteeuwen/go-bindata/...

# 获取根脚本所在目录
sh_path=$(cd `dirname $0`; pwd)

cd $sh_path

echo "cd to $sh_path"

#cur=$(pwd)
#echo "pwd:$cur"
cd ../resources/html/

echo "go-bindata"

go-bindata -pkg=pages -debug -ignore=.idea -o ../../internal/handlers/web/pages/static.go ./...

#cd $cur