FROM golang

ENV GOPROXY https://athens.azurefd.net

RUN go get -u -v github.com/astaxie/beego
RUN go get -u -v github.com/beego/bee


# ADD . .
VOLUME "src/a_beego"

WORKDIR $GOPATH/src/a_beego
RUN ls $GOPATH/src/a_beego
EXPOSE 8080

# RUN go build main.go

# ENTRYPOINT ["./main"]
CMD ["bee","run"]
