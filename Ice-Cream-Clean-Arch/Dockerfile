FROM golang:1.20 AS builder

WORKDIR /opt/app-root

RUN go env -w GOOS=linux GOARCH=amd64 
    
ADD . /opt/app-root
RUN go build -o /opt/app-root/bin/icecream-app cmd/ice-cream-app/main.go 
 
FROM debian:stable-slim
RUN mkdir /opt/app
WORKDIR /opt/app
 
COPY --from=builder /opt/app-root/bin/icecream-app .

EXPOSE 8080

CMD ["./icecream-app"]