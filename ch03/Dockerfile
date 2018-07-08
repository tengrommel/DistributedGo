FROM golang:1.10
ADD . /go/src/hello
RUN go install hello
ENV NAME TENG
ENTRYPOINT /go/bin/hello

# ADD区别与COPY ADD可以添加网络地址
# ENTRYPOINT，表示镜像在初始化时需要执行的命令，不可被重写覆盖，需谨记