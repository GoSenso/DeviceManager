FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN GOARCH=arm GOARM=7 go build ./cmd/DeviceManager/main.go -o ./myapp

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/myapp .
EXPOSE 8080
CMD ["./myapp"]
