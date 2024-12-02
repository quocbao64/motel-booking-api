FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

######## Start a new stage from scratch #######

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /app/main .

EXPOSE 3006
CMD ["./main"]