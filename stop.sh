#!/bin/bash

#杀掉存在的进程
array=(8880 8881 8882 8883 8884)
for i in ${array[@]}
do
    port=$i
    #根据端口号查询对应的pid
    #pid=$(lsof -i:$port |awk '{print $2}' | tail -n 2);
    pid=`lsof -t -i:$port`
    #杀掉对应的进程，如果pid不存在，则不执行
    echo $pid
    if [  -n  "$pid"  ];  then
        kill  -9  $pid;
        echo "kill $port";
    fi
done
