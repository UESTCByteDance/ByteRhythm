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
go mod tidy
```
3.数据库配置

打开`config.ini`，修改以下内容：
```go
DBHost = 127.0.0.1
DBPort = 3306
DBUser = root
DBPassWord =
DBName = tiktok
```
确保你的`Ubuntu20.04`已经装了`MySQL`，并且能够连接上，然后新建数据库`tiktok`

4.配置`ffmpeg`环境

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
5.启动etcd

如果未安装，前往官方网站:<https://github.com/etcd-io/etcd/releases/tag/v3.5.9>下载适合你系统的安装包，解压后找到etcd文件，在对应终端执行：

```bash
./etcd
```
如果权限不够，可以使用`chmod +x etcd`赋予可执行权限再执行`./etcd`。
可以安装`etcdkeeper`进入UI界面进行查看。

6.启动Jaeger
如果未安装，前往官方网站：<https://www.jaegertracing.io/download/>下载适合你系统的安装包，解压后进入，在对应终端执行：
```bash
./jaeger-all-in-one --collector.zipkin.host-port=:9411
```
如果权限不够，可以使用`chmod +x jaeger-all-in-one`赋予可执行权限再执行`./jaeger-all-in-one --collector.zipkin.host-port=:9411`。
可以访问：<http://localhost:16686>进入UI界面。

7.多个终端运行项目（在根目录执行命令）
```bash
go run app/gateway/cmd/main.go
go run app/user/cmd/main.go
go run app/video/cmd/main.go
```

