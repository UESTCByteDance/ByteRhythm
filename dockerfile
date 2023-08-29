#第一阶段
FROM golang:1.17 as builder

WORKDIR /workspace

COPY . .

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
RUN go mod download

RUN go build ./app/gateway/cmd/main.go main.go -o gateway &&\
    go build ./app/user/cmd/main.go main.go -o user &&\
    go run ./app/video/cmd/main.go -o video

#第二阶段
FROM ubuntu:20.04 AS production

## 安装go1.17.8
RUN chmod -R 777 /app/tmp&& cd /app/tmp            
RUN wget https://go.dev/dl/go1.17.8.linux-amd64.tar.gz &&\
    tar -C /usr/local -xzf go1.17.8.linux-amd64.tar.gz &&\
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
    chmod +x etcd &&\
	./etcd

## 安装&运行 Jaeger v3.5.9
RUN wget -c https://github.com/jaegertracing/jaeger/releases/download/v1.48.0/jaeger-1.48.0-linux-amd64.tar.gz &&\
    tar -zxvf jaeger-1.48.0-linux-amd64.tar.gz &&\
	cd jaeger-1.48.0-linux-amd64 &&\
    chmod a+x jaeger-* &&\
    ./jaeger-all-in-one --collector.zipkin.host-port=:9411

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
### 启动 RabbitMQ 服务器
RUN systemctl start rabbitmq-server 
### 配置 RabbitMQ
RUN rabbitmqctl add_user  admin  admin  &&\
    rabbitmqctl set_user_tags admin administrator &&\
    rabbitmqctl  set_permissions -p / admin '.*' '.*' '.*' &&\
    rabbitmqctl status
## 安装&运行 Redis 
RUN apt install redis-server &&systemctl status redis 


WORKDIR /app
COPY --from=builder /workspace/gateway .
COPY --from=builder /workspace/user .
COPY --from=builder /workspace/video .
EXPOSE 8081 8082 8083 8084 8085 8086

RUN chmod +x run.sh
CMD ["./run.sh"]


## docker build -t david945/ByteRhythm:v1.0 .
## docker run -it -p 8081-8086:8081-8086/tcp --name ByteRhythm david945/ByteRhythm:v1.0