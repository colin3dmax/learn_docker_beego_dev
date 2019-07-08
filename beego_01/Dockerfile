FROM golang:1.12 as build

ENV GOPROXY https://athens.azurefd.net
ENV GO111MODULE on

WORKDIR $GOPATH/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /a_beego
ADD . .

EXPOSE 8080

RUN go build main.go && pwd && ls

ENTRYPOINT ["./main"]
# CMD ["main"]
