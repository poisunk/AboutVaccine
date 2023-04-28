FROM golang:alpine
WORKDIR $GOPATH/src/gin_docker
ADD . ./
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
RUN go build -o gin_docker about-vaccine/app
EXPOSE 8080
ENTRYPOINT  ["./gin_docker"]