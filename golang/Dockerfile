FROM golang:1.22 AS builder

WORKDIR /app


COPY . .

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o goapp ./cmd/main.go


FROM alpine:latest


RUN apk --no-cache add ca-certificates

WORKDIR /root/


COPY --from=builder /app/ .



RUN chmod +x ./goapp

EXPOSE 8080

CMD ["./goapp"] 
