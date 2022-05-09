#依赖镜像
FROM golang

#作者信息
MAINTAINER  zay

#工作目录
WORKDIR /app

COPY . .

#在Docker工作目录下执行命令
RUN go build -o main ./main.go

#暴露端口
EXPOSE 8082

#执行项目的命令,上面执行run后的go build会创建一个main的二进制文件
CMD ["/app/main"]