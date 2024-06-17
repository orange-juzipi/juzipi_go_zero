# docker 使用教程

## 命令指南
```bash
# 启动所有服务
# -f 指定docker-compose文件
docker compose -f compose-base.yaml -f compose.yaml up --build

# 启动指定服务
# --build server 表示重新构建服务,server 表示服务名称
# --no-deps 表示不启动依赖服务
docker compose -f compose-base.yaml -f compose.yaml up --build server --no-deps 
```

## 常用命令
```bash
# 启动所有服务
docker compose -f compose-base.yaml -f compose.yaml up --build

# 启动指定服务
docker compose -f compose-base.yaml -f compose.yaml up --build server --no-deps

# 停止所有服务
docker compose -f compose-base.yaml -f compose.yaml down

# 停止指定服务
docker compose -f compose-base.yaml -f compose.yaml down server

# 重启所有服务
docker compose -f compose-base.yaml -f compose.yaml restart

# 重启指定服务
docker compose -f compose-base.yaml -f compose.yaml restart server

# 删除所有服务
docker compose -f compose-base.yaml -f compose.yaml rm

# 删除指定服务
docker compose -f compose-base.yaml -f compose.yaml rm server

# 构建所有服务
docker compose -f compose-base.yaml -f compose.yaml build

# 构建指定服务
docker compose -f compose-base.yaml -f compose.yaml build server

# 构建并启动所有服务
docker compose -f compose-base.yaml -f compose.yaml up --build

# 构建并启动指定服务
docker compose -f compose-base.yaml -f compose.yaml up --build server

# 构建并重启所有服务
docker compose -f compose-base.yaml -f compose.yaml restart

# 构建并重启指定服务
docker compose -f compose-base.yaml -f compose.yaml restart server

# 构建并停止所有服务
docker compose -f compose-base.yaml -f compose.yaml down

# 构建并停止指定服务
docker compose -f compose-base.yaml -f compose.yaml down server

# 构建并删除所有服务
docker compose -f compose-base.yaml -f compose.yaml rm

# 构建并删除指定服务
docker compose -f compose-base.yaml -f compose.yaml rm server
```