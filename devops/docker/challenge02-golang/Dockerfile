FROM golang:alpine as builder

WORKDIR /app
COPY main.go .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/myapp .


FROM scratch

COPY --from=builder /app/myapp .

ENTRYPOINT ["/myapp"]