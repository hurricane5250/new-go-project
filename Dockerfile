# 基础镜像，使用官方的Go镜像
FROM golang:1.24.0 AS builder

# 设置工作目录
WORKDIR /go/src/app

# 设置 Go 模块代理
ENV GOPROXY=https://goproxy.cn,direct

# 复制源代码到容器中
COPY . .

# 安装依赖
RUN go get -d -v ./...

# 构建应用
RUN go build -o /go/bin/app

# 设置环境变量
ENV PATH /go/bin:$PATH

# 暴露指定端口
EXPOSE 8080

# 运行应用
CMD ["app"]