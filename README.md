# ByteRhythm
A repository for minimalist tiktok code

注意：本项目的运行环境为`Ubuntu20.04`，原因是`ffmpeg`对`windows`不够友好,且不支持`Ubuntu20.04`更高的版本，`centos`暂时不清楚是否支持。

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
5.运行项目
```bash
go build
./ByteRhythm
```
也可以通过`bee`工具：
```bash
go install github.com/beego/bee/v2@latest
bee run
```
