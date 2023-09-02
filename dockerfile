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
RUN apt update && apt install -y wget      
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
## 设置时区
RUN apt-get -y update && DEBIAN_FRONTEND="noninteractive" apt -y install tzdata
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /workspace

COPY . .

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
## 安装go 1.20.7
RUN apt update && apt install -y wget  
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
## 安装&运行 etcd v3.5.9
RUN wget https://github.com/etcd-io/etcd/releases/download/v3.5.9/etcd-v3.5.9-linux-amd64.tar.gz &&\
	tar -zxvf etcd-v3.5.9-linux-amd64.tar.gz &&\
	cd etcd-v3.5.9-linux-amd64 &&\
    chmod +x etcd && ln -s ./* /usr/bin/ &&\
	cd ../ &&\
    nohup ./etcd & 

## 安装&运行 Jaeger v3.5.9
RUN wget -c https://github.com/etcd-io/etcd/releases/download/v3.5.9/etcd-v3.5.9-linux-amd64.tar.gz &&\
    tar -zxvf jaeger-1.48.0-linux-amd64.tar.gz &&\
	cd jaeger-1.48.0-linux-amd64 &&\
    chmod a+x jaeger-* && ln -s ./* /usr/bin/ &&\ 
    cd ../ &&\
    nohup ./jaeger-all-in-one --collector.zipkin.host-port=:9411 &

## 安装&运行 RabbitMQ 
### 导入 RabbitMQ 的存储库密钥
RUN apt-get update &&\
    wget -O- https://github.com/rabbitmq/signing-keys/releases/download/2.0/rabbitmq-release-signing-key.asc | sudo apt-key add -
### 将存储库添加到系统
RUN apt-get install apt-transport-https &&\
    cho "deb https://dl.bintray.com/rabbitmq-erlang/debian focal erlang" | sudo tee /etc/apt/sources.list.d/bintray.erlang.list &&\
    echo "deb https://dl.bintray.com/rabbitmq/debian focal main" | sudo tee /etc/apt/sources.list.d/bintray.rabbitmq.list 
### 安装 RabbitMQ 和 Erlang
RUN apt-get install rabbitmq-server
### 启动 RabbitMQ 服务器n 
RUN systemctl start rabbitmq-server 
### 配置 RabbitMQ
# RUN rabbitmqctl add_user  admin  admin  &&\
#     rabbitmqctl set_user_tags admin administrator &&\
#     rabbitmqctl  set_permissions -p / admin '.*' '.*' '.*' &&\
#     rabbitmqctl status
## 安装&运行 Redis 
RUN apt install -y redis-server &&systemctl start redis-server
## 安装&运行 mysql
RUN apt install -y mysql-server &&systemctl start mysql-server
#设置免密登录
ENV MYSQL_ALLOW_EMPTY_PASSWORD yes
RUN systemctl start mysql &&\
    chmod +x mysql_setup.sh && ./mysql_setup.sh &&\
    systemctl restart mysql

WORKDIR /app
COPY --from=builder /workspace/gateway .
COPY --from=builder /workspace/user .
COPY --from=builder /workspace/video .
COPY --from=builder /workspace/relation .
COPY --from=builder /workspace/favorite .
COPY --from=builder /workspace/comment .
EXPOSE 8081 8082 8083 8084 8085 8086

RUN chmod +x run.sh
CMD ["./run.sh"]


## docker build -t david945/byterhythm:v1.0 .
## docker run -it -p 8080-8086:8080-8086/tcp --name byterhythm david945/byterhythm:v1.0