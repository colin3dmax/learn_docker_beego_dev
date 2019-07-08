FROM my_golang

# ADD . .
VOLUME "src/beego_03"

WORKDIR $GOPATH/src/beego_03
RUN ls $GOPATH/src/beego_03
EXPOSE 8080

# RUN go build main.go

# ENTRYPOINT ["./main"]
CMD ["bee","run"]
