#!/bin/bash

# trap 命令用于在 shell 脚本退出时，删掉临时文件，结束子进程。
trap "rm sever;kill 0" EXIT

go build -o sever
./sever -port=8001 &
./sever -port=8002 &
./sever -port=8003 -api=1 &

sleep 2
echo ">>> start test"
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &

wait
