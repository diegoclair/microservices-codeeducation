FROM golang:latest

# Add Maintainer info
LABEL maintainer="Rodrigues Diego <diego93rodrigues@gmail.com>"

# Set the current working directory inside the container 
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container 
COPY . /app

# Add docker-compose-wait tool -------------------
ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

EXPOSE 80

#This is used to run the application with live reload
RUN go get github.com/githubnemo/CompileDaemon

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o myapp

ENTRYPOINT /wait && CompileDaemon --command="./myapp rest"

