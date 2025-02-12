#!/bin/bash

BINARY_NAME="jhzy"  # 比如：your_service_name


# 1. 编译 Go 程序
echo "正在编译 Go 程序..."
go build -o $BINARY_NAME main.go

# 2. 查找并杀掉旧进程
PID=$(ps aux | grep "$BINARY_NAME" | grep -v "grep" | awk '{print $2}')

if [ -n "$PID" ]; then
  echo "找到现有进程 $PID, 正在停止..."
  kill -9 $PID
  sleep 2
  echo "进程已停止"
else
  echo "未找到现有进程"
fi

# 3. 启动新的服务进程
echo "启动新的服务..."
nohup ./$BINARY_NAME > /dev/null 2>&1 &

echo "服务已在后台运行"