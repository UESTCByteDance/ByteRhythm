# ByteRhythm
A repository for minimalist tiktok code

注意：本项目的运行环境为`Ubuntu20.04`，原因是`ffmpeg`对`windows`不够友好,且不支持`Ubuntu20.04`更高的版本，`centos`暂时不清楚是否支持。

如果Ubuntu上没有Golang开发环境，可参考这篇文章进行配置：<https://blog.csdn.net/m0_63230155/article/details/132246694?spm=1001.2014.3001.5502>

# 使用说明
1.克隆到本地
```bash
git clone https://github.com/UESTCByteDance/ByteRhythm.git
```
2.安装依赖
```bash
go mod init ByteRhythm
go mod tidy
```
3.数据库配置

打开`app.conf`，修改以下内容：
```go
username = root
password = 123456
dbHost = 127.0.0.1
dbPort = 3306
database = tiktok
```
确保你的`Ubuntu20.04`已经装了`MySQL`，然后新建数据库`tiktok`

4.配置`ffmpeg`环境

打开终端，依次执行下列命令(建议逐条执行）：
```bash
sudo apt-get -y install autoconf automake build-essential libass-dev libfreetype6-dev libsdl1.2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev pkg-config texi2html zlib1g-dev

sudo apt install -y libavdevice-dev libavfilter-dev libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev

sudo apt-get install yasm

export FFMPEG_ROOT=$HOME/ffmpeg
export CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter"
export CGO_CFLAGS="-I$FFMPEG_ROOT/include"
export LD_LIBRARY_PATH=$HOME/ffmpeg/lib
```
5.配置redis

```bash
1. 更新软件包索引列表。打开终端并使用如下命令：
sudo apt update
2. 安装 Redis 依赖项。使用如下命令：
sudo apt install build-essential tcl
3. 下载最新版本的 Redis。可以从 Redis 的官方网站获取最新版本的 Redis：
wget http://download.redis.io/redis-stable.tar.gz
4. 解压 Redis 压缩包。使用如下命令：
tar xzf redis-stable.tar.gz
5. 进入 Redis 目录。使用如下命令：
cd redis-stable
6. 编译 Redis。使用如下命令：
make
7. 安装 Redis。使用如下命令：
sudo make install
8. 启动 Redis 服务。使用如下命令：
redis-server
现在 Redis 服务已经成功安装并运行在本地机器上。
```

6.运行项目

```bash
go build
./ByteRhythm
```
也可以通过`bee`工具：
```bash
go install github.com/beego/bee/v2@latest
bee run
```
