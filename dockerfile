#第一阶段
FROM ubuntu:20.04 as builder
## 设置时区
RUN apt-get -y update && DEBIAN_FRONTEND="noninteractive" apt -y install tzdata
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /workspace

COPY . .

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

## 安装go1.20.7  
RUN apt install -y wget      
RUN wget https://go.dev/dl/go1.20.7.linux-amd64.tar.gz &&\
    tar -C /usr/local -xzf go1.20.7.linux-amd64.tar.gz &&\
    ## 软链接
    ln -s /usr/local/go/bin/* /usr/bin/
# 配置ffmpeg环境
RUN apt-get  install -y autoconf automake build-essential libass-dev libfreetype6-dev libsdl1.2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev pkg-config texi2html zlib1g-dev
RUN apt install -y libavdevice-dev libavfilter-dev libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev
RUN apt-get install -y yasm
    # 设置环境变量
ENV FFMPEG_ROOT=$HOME/ffmpeg \
    CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter" \
    CGO_CFLAGS="-I$FFMPEG_ROOT/include" \
    LD_LIBRARY_PATH=$HOME/ffmpeg/lib


RUN go mod download

RUN chmod +x build.sh && ./build.sh

#第二阶段
FROM ubuntu:20.04 AS production
WORKDIR /app

## 设置时区
RUN apt-get -y update && DEBIAN_FRONTEND="noninteractive" apt -y install tzdata
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY ./*.sh  .

#本机目录不能使用绝对路径，因为它本身就是一个相对路径
#只会复制本机的config目录下所有文件，而不会创建config目录，所以后面需要指定
COPY ./config ./config 
# COPY . .
RUN mkdir -p ./app/video/tmp/
RUN apt install -y wget systemctl

# 配置ffmpeg环境
RUN apt-get  install -y autoconf automake build-essential libass-dev libfreetype6-dev libsdl1.2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev pkg-config texi2html zlib1g-dev
RUN apt install -y libavdevice-dev libavfilter-dev libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev
RUN apt-get install -y yasm
    # 设置环境变量
ENV FFMPEG_ROOT=$HOME/ffmpeg \
    CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter" \
    CGO_CFLAGS="-I$FFMPEG_ROOT/include" \
    LD_LIBRARY_PATH=$HOME/ffmpeg/lib

## 安装 etcd v3.5.9
RUN wget https://github.com/etcd-io/etcd/releases/download/v3.5.9/etcd-v3.5.9-linux-amd64.tar.gz &&\
    tar -zxvf etcd-v3.5.9-linux-amd64.tar.gz &&\
    cd etcd-v3.5.9-linux-amd64 &&\
    chmod +x etcd  &&\
    mv ./etcd* /usr/local/bin/


## 安装 Jaeger v3.5.9
RUN wget -c https://github.com/jaegertracing/jaeger/releases/download/v1.48.0/jaeger-1.48.0-linux-amd64.tar.gz &&\
    tar -zxvf jaeger-1.48.0-linux-amd64.tar.gz &&\
	cd jaeger-1.48.0-linux-amd64 &&\
    chmod a+x jaeger-* &&\ 
    mv ./jaeger-* /usr/local/bin/
    # nohup ./jaeger-all-in-one --collector.zipkin.host-port=:9411 &

## 安装 RabbitMQ 
### 导入 RabbitMQ 的存储库密钥
RUN wget -O- https://github.com/rabbitmq/signing-keys/releases/download/2.0/rabbitmq-release-signing-key.asc | apt-key add -
### 将存储库添加到系统
RUN apt-get install -y apt-transport-https &&\
    cho "deb https://dl.bintray.com/rabbitmq-erlang/debian focal erlang" | tee /etc/apt/sources.list.d/bintray.erlang.list &&\
    echo "deb https://dl.bintray.com/rabbitmq/debian focal main" | tee /etc/apt/sources.list.d/bintray.rabbitmq.list 
### 安装 RabbitMQ 和 Erlang
RUN apt-get install -y rabbitmq-server

## 安装 注意：Redis安装会自动启动
RUN apt install -y redis-server 


COPY --from=builder /workspace/gateway .
COPY --from=builder /workspace/user .
COPY --from=builder /workspace/video .
COPY --from=builder /workspace/relation .
COPY --from=builder /workspace/favorite .
COPY --from=builder /workspace/comment .
COPY --from=builder /workspace/message .
EXPOSE 8080 16686

# RUN chmod +x /app/run.sh 等效下面语句
RUN cd /app &&chmod +x start.sh
CMD ["/app/start.sh"]


## docker build -t david945/byterhythm:v2.1 .
## docker run -it -p 8080:8080/tcp -p 16686:16686/tcp --name byterhythm david945/byterhythm:v2.1
