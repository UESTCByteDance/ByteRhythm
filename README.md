<h1 align="center" style="border-bottom: none;">ByteRhythm</h1>
<h4 align="center">本项目利用 Golang 以及相关技术如 Gorm、MySQL、Redis、JWT、RabbitMQ、Hystrix、七牛云 等构建了基于 Gin 和 Go-micro的微服务应用，实现了视频处理、对象存储、限流、降级熔断、负载均衡等功能，并通过 Opentracing、Jaeger 等工具进行监控与追踪，Testify进行单元测试，Docker进行容器化部署，形成高可用高性能的分布式服务。</h4>
<div class="labels" align="center">
    <a href="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg">
      <img src="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg" alt="semantic-release">
    </a>
    <a href="https://pkg.go.dev/github.com/UESTCByteDance/ByteRhythm/v2">
      <img src="https://godoc.org/github.com/UESTCByteDance/ByteRhythm?status.svg" alt="Godoc">
    </a>
    <a href="https://github.com/UESTCByteDance/ByteRhythm/blob/master/LICENSE">
      <img src="https://img.shields.io/github/license/UESTCByteDance/ByteRhythm?style=flat-square" alt="license">
    </a>
    <a href="https://github.com/UESTCByteDance/ByteRhythm/issues">
      <img src="https://img.shields.io/github/issues/UESTCByteDance/ByteRhythm?style=flat-square" alt="GitHub issues">
    </a>
    <a href="#">
      <img src="https://img.shields.io/github/stars/UESTCByteDance/ByteRhythm?style=flat-square" alt="GitHub stars">
    </a>
    <a href="https://github.com/UESTCByteDance/ByteRhythm/network">
      <img src="https://img.shields.io/github/forks/UESTCByteDance/ByteRhythm?style=flat-square" alt="GitHub forks">
    </a>
    <a href="https://github.com/UESTCByteDance/ByteRhythm/releases/latest">
      <img src="https://img.shields.io/github/release/UESTCByteDance/ByteRhythm.svg" alt="Release">
    </a>
    <span>
      <img src="https://img.shields.io/github/languages/top/UESTCByteDance/ByteRhythm" alt="language">
    </span>
    <span>
      <img src="https://img.shields.io/github/last-commit/UESTCByteDance/ByteRhythm" alt="last commit">
    </span>
   <span>
      <img src="https://komarev.com/ghpvc/?username=ByteRhythm&label=Views&color=0e75b6&style=flat" alt="访问量统计" />
    </span>
</div>


#### 📚文档地址：<https://uestcbytedance.github.io/>

#### 🥽视频地址：<https://github.com/UESTCByteDance/UESTCByteDance.github.io/blob/main/tiktok.mp4>

#### 😁服务地址：<http://47.113.223.86:8080>

# 使用说明
本项目有v1、v2两个版本，可前往Releases：<https://github.com/UESTCByteDance/ByteRhythm/releases>下载使用，前者是传统的单体架构，用beego实现，后者是微服务架构，由gin+go-micro实现。

下面介绍v2版的使用：

如果不使用docker进行容器化部署，可以参考以下步骤进行本地部署。建议使用环境为`Ubuntu20.04`。

## 1.克隆到本地

```bash
git clone https://github.com/UESTCByteDance/ByteRhythm.git
```

## 2.安装依赖

```bash
go mod tidy
```

## 3.数据库配置

打开`config.ini`，修改以下内容：

```ini
DBHost = 127.0.0.1
DBPort = 3306
DBUser = root
DBPassWord = 123456
DBName = tiktok
```

确保你的`Ubuntu20.04`已经装了`MySQL`，并且能够连接上，然后新建数据库`tiktok`

## 4.配置`ffmpeg`环境

打开终端，依次执行下列命令(逐条执行）：

```bash
sudo apt-get -y install autoconf automake build-essential libass-dev libfreetype6-dev libsdl1.2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev pkg-config texi2html zlib1g-dev

sudo apt install -y libavdevice-dev libavfilter-dev libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev

sudo apt-get install yasm

export FFMPEG_ROOT=$HOME/ffmpeg
export CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter"
export CGO_CFLAGS="-I$FFMPEG_ROOT/include"
export LD_LIBRARY_PATH=$HOME/ffmpeg/lib
```

## 5.启动etcd

如果未安装，前往官方网站:<https://github.com/etcd-io/etcd/releases/tag/v3.5.9>下载适合你系统的安装包并解压。

按需修改配置：

```ini
EtcdHost = 127.0.0.1
EtcdPort = 2379
```

在对应终端执行：

```bash
./etcd
```

如果权限不够，可以使用`chmod +x etcd`赋予可执行权限再执行`./etcd`。
可以安装`etcdkeeper`进入UI界面进行查看。

## 6.启动Jaeger

如果未安装，前往官方网站：<https://www.jaegertracing.io/download/>下载适合你系统的安装包并解压。

按需修改配置：

```ini
JaegerHost = 127.0.0.1
JaegerPort = 6831
```

在对应终端执行：

```bash
./jaeger-all-in-one --collector.zipkin.host-port=:9411
```

如果权限不够，可以使用`chmod +x jaeger-all-in-one`
赋予可执行权限再执行`./jaeger-all-in-one --collector.zipkin.host-port=:9411`。
可以访问：<http://localhost:16686>进入UI界面。

## 7.启动RabbitMQ

如果未安装，前往官方网站：<https://www.rabbitmq.com/install-debian.html>下载安装。

按需修改配置：

```ini
RabbitMQ = amqp
RabbitMQHost = 127.0.0.1
RabbitMQPort = 5672
RabbitMQUser = guest
RabbitMQPassWord = guest
```

确保RabbitMQ能在本地运行。

## 8.配置Redis

如果未安装，打开终端，依次执行下列命令(逐条执行）：

```bash
sudo apt update
sudo apt install redis-server
```

按需修改配置：

```ini
RedisHost = 127.0.0.1
RedisPort = 6379
```

确保Redis能在本地运行。

## 9.配置七牛云

根据你的七牛云账户信息，修改以下配置：

```ini
Bucket = your bucket
AccessKey = your access key
SecretKey =  your secret key
Domain = your domain
```
## 10.运行项目

```bash
//构建项目
chmod +x build.sh
./build.sh

//运行项目                          
chmod +x run.sh
./run.sh                            
```
## 11.运行测试
```bash
cd test
go test -v
```
注：测试文件的参数可能会需要根据实际情况更改。
# docker 运行
## 1.拉取 mysql 镜像并运行
```sh
docker run -d -p 3306:3306 --name tiktok-mysql -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=tiktok mysql/mysql-server:latest
```
## 2.拉取 byterhythm:v2.1 镜像并运行
```sh
docker run -it -p 8080-8086:8080-8086/tcp --name byterhythm david945/byterhythm:v2.1
```


