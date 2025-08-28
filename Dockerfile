FROM registry.cn-shanghai.aliyuncs.com/wangjian3306/golang:1.21.13-alpine AS builder
LABEL authors="wangjian3306"

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/build
WORKDIR /build

# 安装 swag 工具
RUN go install github.com/swaggo/swag/cmd/swag@latest

# 复制项目中的 go.mod 和 go.sum 文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 生成 swagger 文档并将代码编译成二进制可执行文件 app
RUN swag init --parseDependency --parseInternal --parseDepth 1 && go build -o app .


# 接下来创建一个小镜像
FROM scratch

COPY ./static /static
COPY ./conf /conf

# 从builder镜像中把 app 拷贝到根目录
COPY --from=builder /build/app /app

# 声明服务端口
EXPOSE 10086

# 需要运行的命令
ENTRYPOINT ["/app", "-f", "conf/config.yaml"]