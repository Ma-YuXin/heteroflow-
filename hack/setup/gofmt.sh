# 格式化所有go代码文件

find ../ -iname "*.go" -type f -exec gofmt -w {} \; 