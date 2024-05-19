#!/bin/bash

declare -A visitedLibraries # 声明一个关联数组来存储已访问的库

# A function to print the shared library dependencies of a file
print_dependencies() {
    local file="$1"
    local indent="$2"

    # 如果文件已经处理过，就不再处理
    if [[ ${visitedLibraries["$file"]} ]]; then
        return
    fi

    # 标记文件为已处理
    visitedLibraries["$file"]=1

    # 使用ldd命令查找动态链接库，并解析结果
    local libs=$(ldd "$file" 2>/dev/null | grep "=> /" | awk '{print $(NF-1)}')
    for lib in $libs; do
        echo "$indent$lib"
        print_dependencies "$lib" "    $indent"
    done
}
# /mnt/data/nfs/myx/tmp/app/blender-4.1.1-linux-x64/lib/libOpenImageDenoise_device_cuda.so.2.2.2
# Initial call to the function with the target binary and no indentation
print_dependencies "$1" ""