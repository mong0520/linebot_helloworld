
#build stage
FROM golang:1.12 AS builder
WORKDIR /go/src/app
COPY . .
ENV GO111MODULE on
RUN go build
RUN go install -v ./...

#final stage
FROM ubuntu:18.04
RUN apt-get update
#necessary package to make ngrok work
RUN apt-get install -y ca-certificates
COPY --from=builder /go/bin/linebot_helloworld /app
COPY --from=builder /go/src/app/.env /.env
ENTRYPOINT /app
EXPOSE 2000
