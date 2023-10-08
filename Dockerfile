FROM golang:latest

RUN go env -w GO111MODULE=on

RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /home/youtuantuan_societyservice

ADD . /home/youtuantuan_societyservice

RUN go build .

EXPOSE 8080

ENTRYPOINT ["./youtuantuan_societyservice"]



