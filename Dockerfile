FROM golang:alpine as builder

RUN apk update && apk upgrade
RUN apk add musl-dbg gcc libc-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 
COPY . .

RUN go build -tags musl -o ./main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

EXPOSE 8080

CMD ["./main"]
