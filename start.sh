#!/bin/bash

#检测端口是否被占用
array=(8880 8881 8882 8883 8884)
mark=0
for i in ${array[@]}
do
    port=$i
    #根据端口号查询对应的pid
    pid=`lsof -t -i:$port`
    if [  -n  "$pid"  ];  then
        echo "$port"
        mark=1
    fi
done

if [ $mark -eq 1 ]; then
    echo "端口被占用，请先停止进程！"
    exit
fi


cd ./backend
pwd
#编译main文件
if [ -f "./main" ]; then
    echo "main文件存在"
else
    go build ./main.go
fi


./main controller 1 &
sleep 30s
./main hospital 1 &
./main store 1 &
./main store 2 &
./main store 3 &