FROM golang:1.17

ENV SRC_DIR=/go/src/github.com/
WORKDIR $SRC_DIR

RUN apt-get update
RUN go get "github.com/astaxie/beego/migration"
RUN go get "github.com/astaxie/beego/orm"
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/beego/bee"
RUN go get "github.com/astaxie/beego"

ENV GO111MODULE=on