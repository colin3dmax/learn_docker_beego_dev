FROM golang

ENV GOPROXY https://athens.azurefd.net

RUN go get -u -v github.com/astaxie/beego
RUN go get -u -v github.com/beego/bee

# ADD . .
VOLUME "src/beego_02"

WORKDIR $GOPATH/src/beego_02
RUN ls $GOPATH/src/beego_02
EXPOSE 8080

# RUN go build main.go

# ENTRYPOINT ["./main"]
CMD ["bee","run"]
