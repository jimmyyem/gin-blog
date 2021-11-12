FROM golang:latest
ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/gin-blog
COPY . $GOPATH/src/gin-blog
RUN go build -o main && ./main
EXPOSE 8888
ENTRYPOINT ["./main"]