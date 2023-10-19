FROM golang:latest

RUN go env -w GO111MODULE=on

RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /home/ytt

ADD . /home/ytt

RUN go build -o main .

EXPOSE 8080

ENTRYPOINT ["./main"]



