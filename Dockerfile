# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.16 as builder

# 启用go module
# ENV GO111MODULE=on \
#     GOPROXY=https://goproxy.cn,direct

#
# RUN cd ./demo_server

WORKDIR /data/app/go-project

# COPY . .
COPY ./go-project ./

# CGO_ENABLED禁用cgo 然后指定OS等，并go build
#RUN go mod tidy
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o go-project .

# 将这些文件放到了publish文件夹
# RUN mkdir publish && mkdir publish/conf && cp ./go-project publish &&\
#     cp -r ./conf/app.ini publish/conf

# 运行阶段指定基础镜像
FROM centos

WORKDIR /data/app/go-project

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /data/app/go-project .

#
# RUN mkdir -p /app/log

# 指定运行时环境变量
ENV GIN_MODE=debug

EXPOSE 6018

CMD ["./go-project"]