FROM golang:1.16 as builder

WORKDIR /app

# 将当前目录（dockerfile所在目录）下所有文件都拷贝到工作目录下
COPY . /app/

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go env
RUN go mod tidy
RUN go build -o go-cms .

FROM alpine:latest
LABEL MAINTAINER="10512203@qq.com"


#定义时区参数
ENV TZ=Asia/Shanghai 
#设置时区
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo '$TZ' > /etc/timezone 
#设置编码
ENV LANG C.UTF-8

WORKDIR /app

# 将构建产物/app/main拷贝到运行时的工作目录中

COPY --from=builder /app/go-cms /app/config.docker.yaml /app/res /app/res /app/docs /app/
# 将dist文件中的内容复制到 /usr/share/nginx/html/ 这个目录下面
#COPY dist/  /usr/share/nginx/html/

#EXPOSE 40050
# 执行启动命令

CMD ["./go-cms","-c","config.docker.yaml"]

#ENTRYPOINT ./go-cms -c config.docker.yaml
