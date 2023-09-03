#!/bin/bash
cd /usr/local/bin/ &&./etcd &
cd /usr/local/bin/ &&./jaeger-all-in-one --collector.zipkin.host-port=:9411 &
systemctl start rabbitmq-server
systemctl start redis-server

cd /bin/
./gateway &
./user &
./video &
./favorite &
./comment &
./relation &
./message &
