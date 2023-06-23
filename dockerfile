FROM golang:alpine
WORKDIR $GOPATH/src/Vax
ADD . ./
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
RUN go build -o vax vax/app
EXPOSE 8080
RUN chmod 777 ./entrypoint.sh
RUN apk --no-cache add bash
ENTRYPOINT  ["./entrypoint.sh"]