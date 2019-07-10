FROM golang:1.12 as build

# ENV GOPROXY https://athens.azurefd.net
# ENV GO111MODULE on

WORKDIR /beego_00
ADD . .

EXPOSE 8080

RUN go build -mod vendor -o main

ENTRYPOINT ["./main"]
