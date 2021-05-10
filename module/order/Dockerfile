#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./module/order/

#final stage
FROM alpine:latest
ENV SERVER_HTTPPORT=80 DATABASE_HOST=mysql:3306
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app ./app
ADD ./configs/docker.yml ./configs/config.yml
ENTRYPOINT ["./app", "-config", "./configs"]
EXPOSE 80
