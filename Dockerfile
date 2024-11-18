FROM golang:latest
LABEL authors="bianpeijiang"

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/bianpeijiang/ginblog
COPY . $GOPATH/src/github.com/bianpeijiang/ginblog
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./ginblog"]