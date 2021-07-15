
# 编译

1. 安装 `golang`：
>MacOS 下载地址：https://studygolang.com/dl/golang/go1.16.6.darwin-amd64.pkg
> 
>Linux 下载地址：https://studygolang.com/dl/golang/go1.16.6.linux-amd64.tar.gz 

2. 安装docker：`sudo apt install docker.io`。
3. 编译二进制：
```shell
go build -o ./clickhouse-operator cmd/operator/main.go
```
4. Dockerfile路径：
```shell
~/clickhouse-operator$ find . -name "Dockerfile"
./vendor/golang.org/x/net/http2/Dockerfile
./dockerfile/operator/Dockerfile
./dockerfile/metrics-exporter/Dockerfile
```   
5. 编译 docker image(比如是0.14.0版本)：
```shell
# 编译 dbkernel/clickhouse-operator
sudo docker build -t dbkernel/clickhouse-operator:0.14.0 -f ./dockerfile/operator/Dockerfile ./

# 编译 dbkernel/clickhouse-metrics-operator
sudo docker build -t dbkernel/clickhouse-metrics-operator:0.14.0 -f ./dockerfile/metrics-exporter/Dockerfile ./
```
6. 编译成功后，通过`sudo docker images`查询：
```shell
$ sudo docker images
REPOSITORY                             TAG                 IMAGE ID            CREATED             SIZE
dbkernel/clickhouse-metrics-operator   0.14.0              2c8017c5b604        19 minutes ago      50.7MB
dbkernel/clickhouse-operator           0.14.0              e0d0abb3df3b        32 minutes ago      54.2MB
```

>参考 [operator_build_from_sources.md](./docs/operator_build_from_sources.md) 

**注意：**
> 替换路径中的`altinity`为`dbkernel`。

# 推送到 docker hub

1. 登录`docker hub`：
```shell
sudo docker login
# 输入 username、password
```
2. 推送：
```shell
# 用法：docker push username&organizename/dockername:tagname
sudo docker push dbkernel/clickhouse-operator:0.14.0
sudo docker push dbkernel/clickhouse-metrics-operator:0.14.0
```

# 在Kubernetes安装

参考 [operator_installation_details.md](./docs/operator_installation_details.md)


