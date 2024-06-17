# juzipi_go_zero

> 让 go-zero 不再难懂---实战，也是练手项目，从 0 开始，本 issue 和 PR 记录了整个过程，项目虽小，但五脏俱全，前景包括但不限于【微服务、云原生】，同时也会与时俱进，欢迎
> star 和 PR 共同进步。

## 项目地址

- 项目地址：https://github.com/juzi-pi/juzipi_go_zero
- 项目文档：
    - Docker 文档: https://github.com/juzi-pi/juzipi_go_zero/README.Docker.md

## 项目进度

- [ ] 中间件
    - [x] Redis
    - [x] MySQL
    - [ ] JWT
    - [ ] Cache
    - [ ] RabbitMQ
    - [ ] Kafka
    - [ ] ELK
    - [ ] Etcd
    - [ ] 可观测
        - [ ] Prometheus
        - [ ] Grafana
        - [ ] Jaeger
- [ ] 运维部署
    - [x] Docker Compose
    - [ ] k8s
    - [ ] CICD
        - [ ] Jenkens
        - [ ] Drone
        - [ ] Gitea
        - [ ] Gitea Actions

> 本路线会随着项目的进展发生改变，请及时关注。\
> 大家可以把自己的例子放在 Example 目录下并提交，方便大家互相学习。

## 依赖版本

- golang: 1.22.4
- go-zero: 1.6.5
- docker: 26.1.1
- mysql: 8.4.0-1.el9
- redis: 7.2.5

建议直接使用最新版本，勇士不怕有 bug，有问题提 issue，一起解决。

## 快速开始

```shell
# 克隆项目
git clone https://github.com/juzi-pi/juzipi_go_zero.git

# 进入项目目录
cd juzipi_go_zero

# 安装依赖
go mod tidy

# 启动项目
# 加入 --build 参数，会重新构建镜像
# 加入 -d 参数，会后台运行容器 
docker compose -f compose-base.yaml -f compose.yaml up --build

# 访问项目
# 只能访问这两个 URL，其他的会返回 404
http://localhost:8888/from/you
http://localhost:8888/from/me

# 更新某个服务后需要重新打包
docker compose -f compose-base.yaml -f compose.yaml up --build --no-deps --service <服务名>

# 停止项目
docker compose -f compose-base.yaml -f compose.yaml down

```