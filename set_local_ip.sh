#!/bin/bash

# 获取本地 IP 地址
ip=$(ifconfig | grep 'inet ' | grep -v '127.0.0.1' | awk '{print $2}')

# 打印本地 IP 地址
echo "Local IP address: $ip"

# 将 IP 地址设置为全局变量
# 添加 LOCAL_IP 变量到 ~/.bash_profile 文件中
if ! grep -q "export LOCAL_IP" ~/.bash_profile; then
  echo "export LOCAL_IP=$ip" >> ~/.bash_profile
fi
# 加载新的环境变量
source ~/.bash_profile
