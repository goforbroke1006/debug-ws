FROM golang:1.14.0 AS builder
WORKDIR /code/
COPY ./go.mod /code/
COPY ./go.sum /code/
RUN go mod download
COPY ./main.go /code/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o debug-ws .

FROM debian:stretch
WORKDIR /app/
COPY --from=builder /code/debug-ws /app/
COPY ./web/ /app/
ENTRYPOINT [ "debug-ws", "--addr", "0.0.0.0:8080" ]
EXPOSE 8080
