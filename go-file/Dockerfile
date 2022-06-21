FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /hello

# RUN go get -d -v github.com/go-sql-driver/mysql
COPY hello.go .
COPY go.mod .
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello .

FROM alpine:latest as pro、


RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /hello .

CMD ["./hello"]