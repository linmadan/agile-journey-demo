FROM golang:latest
MAINTAINER linmadan <772181827@qq.com>
ENV APP_DIR $GOPATH/src/agile-journey-demo
RUN mkdir -p $APP_DIR
WORKDIR $APP_DIR/
COPY ./pkg pkg
COPY ./conf conf
COPY ./go.mod go.mod
COPY ./main.go main.go
RUN ["ln","-sf","/usr/share/zoneinfo/Asia/Shanghai","/etc/localtime"]
ENV GO111MODULE on
ENV GOPROXY https://goproxy.io
RUN ["go","mod","tidy"]
RUN ["go","build"]
ENTRYPOINT ["./agile-journey-demo"]