FROM golang:1.16

COPY . .

RUN go build -o server server.go
CMD [ "./server" ]